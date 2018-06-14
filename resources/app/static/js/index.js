let index = {
    bindGenerateBtn: function() {
        let generateBtn = document.getElementById("generate-seed-btn")
        generateBtn.onclick = function() {
            astilectron.sendMessage({"name": "generate"}, function(message) {
                if (message.name === "error") {
                    asticode.notifier.error(message.payload);
                    return
                }

                seedContainer = document.getElementById("seed-container")
                seedContainer.innerHTML = message.payload.seed
                console.log(message.payload.seed)

                addressContainer = document.getElementById("address-container")
                addressContainer.innerHTML = message.payload.addresses[0]
                console.log(message.payload.addresses[0])
            })
        }
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            index.bindGenerateBtn();
        })
    }
};
