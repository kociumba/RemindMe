<script>
    import autoAnimate from "@formkit/auto-animate";
    import { LoadSettings, UpdateSettings } from "../wailsjs/go/main/App.js";
    import { onMount, afterUpdate } from "svelte";
    import { writable } from "svelte/store";
    import { Checkbox } from "@svelteuidev/core";
	import { OpenEasterEgg } from "../wailsjs/go/main/App";

    let section = "";
    // let key = "";  // deprecated variable, previously used for qurrying the settings file
    let singeleOutput = "";
    let settingsChanged = "";
    let output = writable([]);

    const settings = ["show_notification", "test_setting"];
    let boolOutput = writable([]);

    onMount(async () => {
        await getSettings();
    });

    async function getSettings() {

        output.set([]);

        for (const setting of settings) {
            singeleOutput = await LoadSettings(section, setting);
            output.update((arr) => [...arr, singeleOutput]);
        }
    }

    async function handleChange(setting, item) {
        const itemReverse = item === "true" ? "false" : "true";

        console.info(setting, itemReverse);
        await UpdateSettings(section, setting, itemReverse);

        await getSettings();
    }

</script>

<main >
    {#each $output as item, index}
        <div class="checkboxContainer">
            <!-- {item}
            | -->
            {settings[index]}
            <Checkbox
                aria-label={settings[index]}
                checked={item === "true"}
                color="orange"
                size="sm"
                style="margin-left: 10px;"
                on:click={() => handleChange(settings[index], item)}
            />
        </div>
    {/each}
    <Checkbox 
        color="orange"
        size="sm"
        style="margin-left: 10px;"
        on:click={OpenEasterEgg("8000")}
    />
</main>

<style>
    .checkboxContainer {
        display: flex;
        flex-direction: row;
        justify-content: center;
        margin: 10px;
    }
</style>
