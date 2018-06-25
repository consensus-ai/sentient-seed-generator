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
            <div class="btn generate-btn" id="generate-seed-btn" onClick={() => seedGenerator.generate()}>Generate Wallet</div>

            <div class="seed-container">
                <div class="label">Backup seed:</div>
                <div class="seed" id="seed">{seedGenerator.data.seed}</div>
                <div class="btn-copy" onClick={copySeed}>copy</div>
                <div class="disclaimer"><b>Warning!</b> save this seed in a safe place offline. If you lose it, your funds will be lost forever.</div>
            </div>

            <div class="address-container">
                <div class="label">Public address:</div>
                <div class="address" id="address">{seedGenerator.data.address}</div>
                <div class="btn-copy" onClick={copyAddress}>copy</div>
            </div>
        </div>
      </div>
    )
  }
}

const copySeed = () => {
    var seed = document.getElementById("seed").textContent;
    const el = document.createElement('textarea');
    el.value = seed;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
}

const copyAddress = () => {
    var seed = document.getElementById("address").textContent;
    const el = document.createElement('textarea');
    el.value = seed;
    document.body.appendChild(el);
    el.select();
    document.execCommand('copy');
    document.body.removeChild(el);
}

// Render top-level component, pass controller data as props
const render = () =>
	preact.render(<SeedGenerator />, document.getElementById('app'), document.getElementById('app').lastElementChild);

// Call global render() when controller changes
seedGenerator.render = render;

// Render of the first time
render();
