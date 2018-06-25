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
                <div class="disclaimer"><b>Warning!</b> save this seed in a safe place offline. If you lose it, your funds will be lost forever.</div>
            </div>

            <div class="address-container">
                <div class="label">Public address:</div>
                <div class="address" id="address">{seedGenerator.data.address}</div>
            </div>
        </div>
      </div>
    )
  }
}

// Render top-level component, pass controller data as props
const render = () =>
	preact.render(<SeedGenerator />, document.getElementById('app'), document.getElementById('app').lastElementChild);

// Call global render() when controller changes
seedGenerator.render = render;

// Render of the first time
render();
