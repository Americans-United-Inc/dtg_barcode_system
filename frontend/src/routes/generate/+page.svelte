<script>
        import { Generate } from "../../lib/wailsjs/go/main/App.js"

        let filename = ""
        let success = false
        let error = ""

        /**
         *
         * @param {string} filename
         */
        function generate(filename) {
            Generate(filename).then((result) => ([success, error] = result))
            document.querySelector(".error").innerHTML(success)
        }
    </script>
    <div class="container audiowide-regular">
        <div class="titles">
            <h1>Enter Filepath to Generate Barcode</h1>
        </div>
        <div class="form-container">
            <!-- svelte-ignore a11y-autofocus -->
            <div class="inputs">
                <label for="filepath">Filepath</label>
                <input type="text" id="filepath" name="filepath" bind:value={filename} autofocus>
            </div>
            <button id="generateButton" class="btn" on:click={generate(filename)}>
                Generate
            </button>
            <div class="error">
                <span>{error}</span>
            </div>
            <a href="../">
            <div class="btn">
                <span>Exit</span>
            </div>
            </a>
        </div>
    </div>
