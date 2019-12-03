if (typeof window.ethereum === 'undefined') {
  alert('Please install metamask to continue!');
    return;
}

const accounts = await ethereum.enable();

// A JSON Serialized version of signedData from our go backend:
const challenge = ...
const data = {
    types: {
      EIP712Domain: challenge.types['EIP712Domain'].types,
    [challenge.primary_type]: challenge.types[challenge.primary_type].types,
  },
  domain: {
    name: challenge.domain.name,
    version: challenge.domain.version,
    salt: challenge.domain.salt,
    chainId: challenge.domain.chain_id,
  },
  primaryType: challenge.primary_type,
  message: {
    nonce: challenge.message.nonce,
    timestamp: challenge.message.timestamp,
    address: challenge.message.address,
  },
};

web3.currentProvider.sendAsync({
    method: 'eth_signTypedData_v3',
  params: [accounts[0], JSON.stringify(data)],
  from: accounts[0],



/////////////////////////////////////////////

const onSignatureComplete = (err, result) => 
    if (err || result.error) {
      throw new Error(err || result.error)
    }
    const signature = result.result.substring(2);
    const r = `0x${signature.substring(0, 64)}`;
    const s = `0x${signature.substring(64, 128)}`;
    const v = parseInt(signature.substring(128, 130), 16);
    
    console.log(`Signature complete:\n\tr: ${r}\n\ts: ${s}\n\tv: ${v}`);
    console.log('Signature length: ', signature.length);

    sendToServerForVerification(signature);
}
