<script>
        import { PrintFile } from "../../lib/wailsjs/go/main/App.js"

        let filepath = ""
        let quantity = 1
        let success = {}


        /**
         * @param {string} filepath
         * @param {number} quantity
         */
        function printFile(filepath, quantity) {
            PrintFile(filepath, quantity).then((result) => (resp = result))
            if (!success) {
                let errors = document.querySelector(".error")
                errors.innerHTML = "<span class='error'>Invalid File</span>"
            }
            document.querySelector('#filepath').value = ""
            document.querySelector('#quantity').value = 1
        }
    </script>


    <div class="container audiowide-regular">
        <h1>Scan a Barcode</h1>
        <div class="form-container">
            <div class="inputs">
                <label for="filepath">Filepath</label>
                <!-- svelte-ignore a11y-autofocus -->
                <input type="text" id="filepath" name="filepath" bind:value={filepath} autofocus>
                <label for="quantity">Quantity</label>
                <input type="number" id="quantity" name="quantity" bind:value={quantity} min="1" max="999">
            </div>
            <div class="errors"></div>
            <div class="buttons">
                <button id="printButton" class="btn audiowide-regular" on:click={printFile(filepath, quantity)}>
                    Print
                </button>
                <a href="../">
                    <div class="btn">
                        <span>Exit</span>
                    </div>
                </a>
            </div>
        </div>
    </div>
