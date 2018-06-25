/** @jsx preact.h */

class SeedGenerator extends preact.Component {
	constructor(props) {
		super(props)
	}
  render() {
    return (
      <div class="body-container">
        <div class="header">
            <div class="logo"></div>
            <span>CONSENSUS</span>
        </div>

        <div class="form-container">
            <div class="btn generate-btn" id="generate-seed-btn" onClick={generateNewSeed}>Generate Wallet</div>

            <div class="seed-container">
                <div class="label">Backup seed:</div>
                <div class="seed" id="seed">{seedGenerator.data.seed}</div>
                <div class="copy-container">
                    <div class="btn btn-copy" onClick={copySeed}>Copy</div>
                    <div class="copy-confirmation hidden" id="seed-copy-confirmation">Seed copied to clipboard!</div>
                </div>
                <div class="disclaimer"><b>Warning!</b> save this seed in a safe place offline. If you lose it, your funds will be lost forever.</div>
            </div>

            <div class="address-container">
                <div class="label">Public address:</div>
                <div class="address" id="address">{seedGenerator.data.address}</div>
                <div class="copy-container">
                    <div class="btn btn-copy" onClick={copyAddress}>Copy</div>
                    <div class="copy-confirmation hidden" id="address-copy-confirmation">Address copied to clipboard!</div>
                </div>
            </div>
        </div>
      </div>
    )
  }
}

const generateNewSeed = () => {
    console.log("generating");
    seedGenerator.generate();
    document.getElementById("seed-copy-confirmation").classList.add("hidden");
    document.getElementById("address-copy-confirmation").classList.add("hidden");
}

const copySeed = () => {
    var seed = document.getElementById("seed").textContent;
    const el = document.createElement('textarea');
    el.value = seed;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);

    document.getElementById("seed-copy-confirmation").classList.remove("hidden");
    document.getElementById("address-copy-confirmation").classList.add("hidden");
}

const copyAddress = () => {
    var address = document.getElementById("address").textContent;
    const el = document.createElement('textarea');
    el.value = address;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);

    document.getElementById("address-copy-confirmation").classList.remove("hidden");
    document.getElementById("seed-copy-confirmation").classList.add("hidden");
}

// Render top-level component, pass controller data as props
const render = () =>
	preact.render(<SeedGenerator />, document.getElementById('app'), document.getElementById('app').lastElementChild);

// Call global render() when controller changes
seedGenerator.render = render;

// Render of the first time
render();
