let index = {
    bindGenerateBtn: function() {
        let generateBtn = document.getElementById("generate-seed-btn")
        generateBtn.onclick = function() { index.generateNewSeed(); }
    },
    generateNewSeed: function() {
        astilectron.sendMessage({"name": "generate"}, function(message) {
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return;
            }

            seedContainer = document.getElementById("seed");
            seedContainer.innerHTML = message.payload.seed;

            addressContainer = document.getElementById("address");
            addressContainer.innerHTML = message.payload.addresses[0];
        })
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            index.bindGenerateBtn();

            index.generateNewSeed();
        })
    }
};
