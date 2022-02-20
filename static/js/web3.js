if (typeof window !== "undefined" && typeof window.ethereum !== "undefined") {
  // We are in the browser and metamask is running.
  window.ethereum.request({ method: "eth_requestAccounts" });
  //web3 = new Web3(window.ethereum);
  
  // detect Metamask account change
  window.ethereum.on('accountsChanged', function (accounts) {
    console.log('accountsChanges',accounts);
  });
} else {
  // We are on the server *OR* the user is not running metamask
  console.log("no MetaMask detected");
  alert("Unfortunately you will not be able to interact with the Ethereum Blockchain as no MetaMask exists with your browser.\n\nPlease try using Google Chrome with the MetaMask extension.");
}