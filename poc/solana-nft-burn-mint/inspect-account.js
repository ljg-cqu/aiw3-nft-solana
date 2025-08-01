require('dotenv').config();
const { Connection, PublicKey } = require('@solana/web3.js');

async function inspectAccount(accountAddress) {
    const solanaNetwork = process.env.SOLANA_NETWORK;
    let endpoint = `https://api.${solanaNetwork}.solana.com`;
    if (solanaNetwork === 'localnet') {
        endpoint = 'http://localhost:8899';
    }

    const connection = new Connection(endpoint, 'confirmed');
    const accountInfo = await connection.getAccountInfo(new PublicKey(accountAddress));

    if (accountInfo === null) {
        console.log('Account not found');
    } else {
        console.log('Account Info:');
        console.log('  Lamports:', accountInfo.lamports);
        console.log('  Owner:', accountInfo.owner.toBase58());
        console.log('  Executable:', accountInfo.executable);
        console.log('  Rent Epoch:', accountInfo.rentEpoch);
        console.log('  Data Length:', accountInfo.data.length);
        //console.log('  Data:', accountInfo.data); // Uncomment to view raw data
    }
}

const accountAddress = process.argv[2];
if (!accountAddress) {
    console.error('Please provide an account address as a command-line argument.');
    process.exit(1);
}

inspectAccount(accountAddress);
