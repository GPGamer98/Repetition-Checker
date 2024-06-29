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
    let amount = 0

    let fileName: string = ""
    
    function select() {
        SelectFile().then((result) => {
            file = result
        })
    }

    async function analyze() {
        if (file !== "") {
            try {
                // @ts-ignore
                countResult = await AnalyzeCSVFiles(file);
                if (countResult) {
                    counted = true
                    amount = countResult.length
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

    function total() {
        amount = 0
        countResult.forEach(element => {
            if (element.Count >= filter) {
                amount += 1
            }
        });
    }
</script>

<div class="top">
    <div class="logo">
        <div class="logo-box"><img src="{logo}" alt="Cogito Ergo Vet"></div>
    </div>
    <div class="title">
        <h1>Contatore di Apparenze</h1>
    </div>
</div>

<div class="input">
    <div class="select">
        <button on:click={select}>Seleziona</button>
        <div class="input-div"><input type="text" bind:value={file}></div>
    </div>
    
    <div class="count-div"><button class="count" on:click={analyze}>Conta</button></div>
</div>

<div class="output">
    {#if counted}
        <div class="top-wr">
            <div class="top-cont">
                        <button class="close-btn" on:click={() => {counted = false; countResult = []}}>Cancella</button>
                        <div class="filter">
                            <span class="material-symbols-outlined">
                                filter_alt
                            </span>
                            Maggiore o uguale a:
                            <input type="number" name="filter" id="filter-inp" class="filter-inp" bind:value={filter} on:input={total}>
                        </div>
                        Totale: {amount}
            </div>
        </div>
    
        <div class="table-container">
            <table class="index" id="index-table">
                <tr>
                    <th class="index" id="index-h">e</th>
                </tr>
    
                {#each countResult as element}
                    {#if element.Count >= filter}
                        <tr class="index" id="index-tr">
                            <td class="index" id="index-td">{countResult.indexOf(element) + 1}</td>
                        </tr>
                    {/if}
                {/each}
            </table>
            <table id="data">
                <tr>
                    <th>Email</th>
                    <th>N° volte</th>
                </tr>
    
                {#each countResult as element}
                    {#if element.Count >= filter}
                        <tr>
                            <td>{element.Entry}</td>
                            <td>{element.Count}</td>
                        </tr>
                    {/if}
                {/each}
            </table>
        </div>
        <div class="export-cont">
            <div class="export">
                <div class="warning">Attenzione! Non esportare il file nella stessa cartella dei file da leggere.<button class="material-symbols-outlined info" on:click={callInfo}>info</button>
                </div>
                <div class="export-btn">
                    <button class="export-button" on:click={exportData}>
                        <span class="material-symbols-outlined">download</span>Esporta CSV</button>
                </div>
            </div>
        </div>
    {:else}
        <div class="waiting">
            <h2>Seleziona una cartella e clicca Conta!</h2>
        </div>
    {/if}
</div>

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
        border-bottom: #161616 4px solid;
    }

    @media (max-width: 60rem) {
        .input, .output {
            background-color: #161616;
            padding: 1rem 0 1rem 0;
            margin: 0 3rem 0 3rem;
            border-radius: 16px;
            filter: drop-shadow(0 8px 5px rgb(15, 15, 15));
        }
    }

    @media (min-width: 60rem) {
        .input, .output {
            background-color: #161616;
            padding: 1rem 0 1rem 0;
            margin: 0 6rem 0 6rem;
            border-radius: 16px;
            filter: drop-shadow(0 8px 5px rgb(15, 15, 15));
        }
    }

    @media (min-width: 90rem) {
        .input, .output {
            background-color: #161616;
            padding: 1rem 0 1rem 0;
            margin: 0 12rem 0 12rem;
            border-radius: 16px;
            filter: drop-shadow(0 8px 5px rgb(15, 15, 15));
        }
    }

    @media (min-width: 110rem) {
        .input, .output {
            background-color: #161616;
            padding: 1rem 0 1rem 0;
            margin: 0 21rem 0 21rem;
            border-radius: 16px;
            filter: drop-shadow(0 8px 5px rgb(15, 15, 15));
        }
    }

    .input {
        padding: 1.5rem 0 1.5rem 0;
    }

    .output {
        margin-top: 3rem;
        margin-bottom: 2rem;
    }

    .select {
        display: flex;
        /* padding: 4rem 1rem 4rem 1rem; */
        justify-content: center;
        align-items: center;
    }

    .count-div {
        margin-top: 2rem;
        display: flex;
        justify-content: center;
        background-color: transparent;
    }

    .count {
        font-size: 2rem;
        padding: 1rem;
        background-color: #0f0f0f;
    }

    /* .count:hover {
        filter: drop-shadow(0 0 16px rgb(185, 185, 185));
    } */

    .waiting {
        display: flex;
        height: 20vh;
        justify-content: center;
        align-items: center;
    }

    .top-cont {
        margin: 0 1rem;
        height: 3rem;
        width: 95vw;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .top-wr {
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
        margin: 1rem 1rem 1.5rem 1rem;
    }

    .index {
        margin: 0;
        width: fit-content;
        padding-right: 0;
        text-align: center;
    }
    
    #index-table {
        margin: 0;
        background-color: #0f0f0f;
        border-radius: 8px 0 0 8px;
        text-align: center;
    }

    #index-h {
        border: none;
        color: transparent;
        width: fit-content;
    }

    #index-td {
        /* border: none; */
        padding-left: 0;
        padding-right: 0;
        width: fit-content;
        text-align: center;
    }

    #index-tr {
        padding: 0;
        width: fit-content;
        text-align: center;
    }

    .export-cont {
        display: flex;
        margin: 0 1rem;
    }

    .export-btn {
        display: flex;
        width: 93vw;
        justify-content: right;
        align-items: center;
    }

    .export {
        width: 100%;
        display: flex;
        justify-content: center;
        margin: .1rem 0 0 0;
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
        background-color: rgb(199, 174, 65);
        padding: .7rem;
        display: flex;
        align-items: center;
        width: 78vw;
    }

    .info {
        background-color: transparent;
        color: black;
        border: none;
        animation: zoom 1s ease-in-out infinite;
    }

    .info:hover {
        color: #494949;
        background-color: transparent;
    }

    img {
        display: flex;
        justify-content: center;
        width: 7rem;
    }

    button {
        background-color: #0f0f0f;
        color: white;
        border: none;
        /* border: 1px solid rgb(34, 34, 34); */
        padding: .5rem;
        border-radius: 16px;
        transition: all linear 100ms;
    }

    button:hover {
        cursor: pointer;
        background-color: rgb(36, 36, 36);
    }

    input {
        width: 48vw;
        margin-left: 1rem;
        color: white;
        background-color: #0f0f0f;
        border: 1px solid rgb(34, 34, 34);
        border-radius: 38px;
        height: 1.5rem;
        padding: 0 .5rem 0 .5rem;
    }

    table {
        border-collapse: collapse;
        width: 80vw;
        border-radius: 0 8px 8px 0;
        background-color: #0f0f0f;
    }

    td, th {
        /* border: 1px white solid; */
        text-align: left;
        padding: 12px;
    }

    span {
        margin-right: 4px;
    }

    input::-webkit-outer-spin-button,
    input::-webkit-inner-spin-button {
        /* display: none; <- Crashes Chrome on hover */
        -webkit-appearance: none;
        margin: 0; /* <-- Apparently some margin are still there even though it's hidden */
    }

    input[type=number] {
        appearance: unset;
        -moz-appearance: textfield; /* Firefox */
    }
</style>
