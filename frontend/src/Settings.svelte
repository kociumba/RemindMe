<script>
	import Settings from './Settings.svelte';
    import autoAnimate from "@formkit/auto-animate";
    import { LoadSettings, UpdateSettings } from "../wailsjs/go/main/App.js";
    import { onMount, afterUpdate } from "svelte";
    import { writable } from "svelte/store";
    import { Checkbox } from "@svelteuidev/core";

    let section = "";
    // let key = "";  // deprecated variable, previously used for qurrying the settings file
    let singeleOutput = "";
    let settingsChanged = "";
    let itemReverse = "";
    let output = writable([]);

    const settings = ["show_notification", "test_setting"];
    let boolOutput = writable([]);

    onMount(async () => {
        await getSettings();
    });

    async function getSettings() {
        for (const setting of settings) {
            singeleOutput = await LoadSettings(section, setting);
            output.update((arr) => [...arr, singeleOutput]);
            console.log(Boolean(output));
        }
    }

    async function handleChange(setting, item) {
        if (item === "true") {
            itemReverse = "false";
        } else if (item === "false") {
            itemReverse = "true";
        }

        console.log("change detected")
        console.log(setting, itemReverse);
        settingsChanged = await UpdateSettings(section, setting, itemReverse);
        console.log(settingsChanged);

        await getSettings();
    }

</script>

<main use:autoAnimate>
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
</main>

<style>
    .checkboxContainer {
        display: flex;
        flex-direction: row;
        justify-content: center;
        margin: 10px;
    }
</style>
