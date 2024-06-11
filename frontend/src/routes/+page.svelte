<script lang="ts">
    import { SelectFile, AnalyzeCSVFiles, ExportDataToCSV, SelectExport, InfoWindow } from "$lib/wailsjs/go/main/App"

    import logo from "$lib/assets/logo.png"

    interface EntryCount {
        Entry: string;
        Count: number;
    }

    let file: string = ""
    let counted = false
    let countResult: EntryCount[];
    let filter = 0

    let fileName: string = ""
    
    function select() {
        SelectFile().then((result) => {
            file = result
        })
    }

    async function analyze() {
        if (file !== "") {
            try {
                countResult = await AnalyzeCSVFiles(file);
                if (countResult) {
                    counted = true
                }
                console.log(countResult.length + 1); // Or further processing/display logic
            } catch (error) {
                console.error("Error:", error);
            }
        }
    }

    function exportData() {
        SelectExport().then((result) => {
            ExportDataToCSV(countResult, result)
        })
    }

    function callInfo() {
        InfoWindow("export")
    }
</script>

<div class="logo">
    <div class="logo-box"><img src="{logo}" alt="Cogito Ergo Vet"></div>
</div>
<div class="title">
    <h1>Contatore di Apparenze</h1>
</div>

<div class="select">
    <button on:click={select}>Seleziona</button>
    <div class="input-div"><input type="text" bind:value={file}></div>
</div>

<div class="count-div"><button class="count" on:click={analyze}>Conta</button></div>

{#if counted}
    <div class="top-wr">
        <div class="top-cont">
                    <button class="close-btn" on:click={() => {counted = false; countResult = []}}>Cancella</button>
                    <div class="filter">
                        <span class="material-symbols-outlined">
                            filter_alt
                        </span>
                        Maggiore o uguale a: 
                        <input type="number" name="filter" id="filter-inp" class="filter-inp" bind:value={filter}>
                    </div>
                    Totale: {countResult.length + 1}
        </div>
    </div>

    <div class="table-container">
        <table>
            <tr>
                <th>Email</th>
                <th>N° volte</th>
            </tr>

            {#each countResult as element}
                <tr>
                    {#if element.Count >= filter}
                        <td>{element.Entry}</td>
                        <td>{element.Count}</td>
                    {/if}
                </tr>
            {/each}

        </table>
    </div>
    <div class="export-cont">
        <div class="export">
            <div class="warning">Attenzione! Non esportare il file nella stessa cartella dei file da leggere.<button class="material-symbols-outlined info" on:click={callInfo}>info</button>
            </div>
            <div class="export-btn">
                <button class="export-button" on:click={exportData}><span class="material-symbols-outlined">
                    download
                    </span>Esporta CSV</button>
            </div>
        </div>
    </div>
{:else}
    <div class="waiting">
        <h2>Seleziona una cartella e clicca Conta!</h2>
    </div>
{/if}

<!-- <footer>
    <div class="footer-div">
        <p>Made with ❤️ by Giacomo for Cogito Ergo Vet</p>
    </div>
</footer> -->

<style>
    @import url('https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300..800;1,300..800&display=swap');
    
    @keyframes zoom {
        0% {
            transform: scale(1);
        }
        50% {
            transform: scale(1.2);
        }
        100% {
            transform: scale(1);
        }
    }

    :root {
        font-family: "Open Sans", sans-serif;
        color: white;
        
    }

    .title {
        display: flex;
        justify-content: center;
        margin-top: 0;
        padding-top: 0;
    }

    .logo {
        margin: 8px;
        padding: 1rem 0 2rem 0;
        display: flex;
        justify-content: center;
        border-bottom: white 1px solid;
    }



    .select {
        display: flex;
        padding: 4rem 1rem 4rem 1rem;
        justify-content: center;
        align-items: center;
    }

    .count-div {
        display: flex;
        justify-content: center;
    }

    .count {
        font-size: 2rem;
        padding: 1rem;
    }

    .count:hover {
        filter: drop-shadow(0 0 16px rgb(185, 185, 185));
    }

    .waiting {
        display: flex;
        height: 20vh;
        justify-content: center;
        align-items: center;
    }

    .top-cont {
        margin-top: 1rem;
        height: 3rem;
        width: 95vw;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .top-wr {
        margin-top: 1.2rem;
        display: flex;
        justify-content: center;
    }

    .filter {
        display: flex;
        align-items: center;
    }

    .filter-inp {
        width: 5rem;
    }

    .table-container {
        display: flex;
        justify-content: center;
        margin: .5rem 0 1rem 0;
    }

    .export-cont {
        display: flex;
        justify-content: center;
    }

    .export-btn {
        display: flex;
        width: 93vw;
        justify-content: right;
        align-items: center;
    }

    .export {
        width: 95vw;
        display: flex;
        justify-content: center;
        margin: .1rem 0 1rem 0;
    }

    .export-button {
        display: flex;
        align-items: center;
        padding-right: 10px;
    }

    .warning {
        color: rgb(44, 44, 44);
        font-size: .9rem;
        border-radius: 8px;
        background-color: rgb(255, 217, 47);
        padding: .7rem;
        display: flex;
        align-items: center;
    }

    .info {
        color: black;
        border: none;
        animation: zoom 1s ease-in-out infinite;
    }

    .info:hover {
        background-color: transparent;
    }

    img {
        display: flex;
        justify-content: center;
        width: 7rem;
    }

    button {
        background-color: transparent;
        color: white;
        padding: .5rem;
        border: 1px solid white;
        border-radius: 48px;
        transition: all linear 100ms;
    }

    button:hover {
        cursor: pointer;
        background-color: rgb(70, 70, 70);
    }

    input {
        width: 48vw;
        margin-left: 1rem;
        color: white;
        background-color: transparent;
        border: 1px solid white;
        border-radius: 38px;
        height: 1.5rem;
        padding: 0 .5rem 0 .5rem;
    }

    table {
        border-collapse: collapse;
        width: 95vw;
        border-radius: 30px;
        border: 1px white solid;
    }

    td, th {
        border: 1px white solid;
        text-align: left;
        padding: 8px;
    }

    span {
        margin-right: 4px;
    }
</style>
