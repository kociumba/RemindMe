<script>
	import autoAnimate from "@formkit/auto-animate";
	import { Modal } from "@svelteuidev/core";
	import { ScheduleNotification } from "../wailsjs/go/main/App";

	let reminderList = [];
	let newReminderText = "";
	let emptyReminderWarning = "";
	let result;
	let opened = false;
	let notificationTime = "";
	let currentReminderTxt;
	let IsNotifScheduled;

	// Add a new reminder to the list
	function addReminder() {
		if (newReminderText == "") {
			emptyReminderWarning = "Please enter a reminder";
			opened = false;
			return;
		} else {
			emptyReminderWarning = "";
			opened = true;
		}

		currentReminderTxt = newReminderText;

		const id = reminderList.length + 1;
		const reminder = { id, text: newReminderText, completed: false };
		reminderList = [...reminderList, reminder];
		newReminderText = ""; // can't remember why this resets but i really don't feel like breaking th app again

	}

	// Delete a reminder from the list
	/**
	 * @param {any} id
	 */
	function deleteReminder(id) {
		console.log(id);
		reminderList = reminderList.filter((reminder) => reminder.id !== id);

	}

	$: {
		if (notificationTime) {
			result = parseTime(notificationTime);
			console.log("parseTime ran: ", notificationTime);
		}
  	}

	$: {
		if (result !== undefined) {
			// Call your desired function when result is defined
			sendNotification(result);

			// Revert the result variable to undefined
			result = undefined;
		}
	}

	/**
     * @param {number} result
     */
	async function sendNotification(result) {
		IsNotifScheduled = await ScheduleNotification(result, currentReminderTxt, notificationTime);
		console.log("sendNotification ran: ", IsNotifScheduled);
		setTimeout(() => {
			const id = reminderList.findIndex((reminder) => reminder.text === currentReminderTxt); // TODO: account for multiple reminders
			deleteReminder(id+1);
		}, result * 1000);
	}

	// parse the notification notificationTime var into hours mins or seconds based on the suffix eg. h, min or s
	/**
     * @param {string} notificationTime
     */
	function parseTime(notificationTime) {
		if (notificationTime.includes("h")) {
			return parseInt(notificationTime.split("h")[0]) * 60 * 60;
		} else if (notificationTime.includes("min")) {
			return parseInt(notificationTime.split("min")[0]) * 60;
		} else if (notificationTime.includes("s")) {
			return parseInt(notificationTime.split("s")[0]);
		} else {
			console.error("parseTime ran: ", notificationTime, "failed"); //TODO: implement the error display on the ui
		}

	}

</script>

<main>
	<div class="container">
		<h1>Reminder App</h1>

		<ul use:autoAnimate>
			{#each reminderList as reminder (reminder.id)}
				<button on:click={() => deleteReminder(reminder.id)}>
					<li class:completed={reminder.completed} class="listButton">
						{reminder.text}
					</li>
				</button>
			{/each}
		</ul>

		<input
			type="text"
			placeholder="Enter a reminder"
			bind:value={newReminderText}
			on:keydown={(e) => e.key === "Enter" && addReminder()}
		/>

		<div class="addButtonContainer">
			<button on:click={addReminder} class="addButton">
				<svg
					width="24"
					height="24"
					viewBox="0 0 24 24"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
				>
					<path
						d="M6 12H18M12 6V18"
						stroke="white"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
					/>
				</svg>
			</button>
		</div>

		<h4>{emptyReminderWarning}</h4>

		<Modal
			centered
			{opened}
			on:close={() => (opened = false)}
			withCloseButton={false}
			target={'main'}
			closeOnEscape={true}
			closeOnClickOutside={true}
			style="" 
		> 											<!-- TODO: make this look half decent  -->
			remind me in {notificationTime}

			<input
				type="text"
				placeholder="Enter a reminder"
				bind:value={notificationTime}
			/>

		</Modal>
	</div>
</main>

<style>
	.container {
		text-align: center;
		padding: 1em;
		max-width: 400px;
		margin: 0 auto;
	}

	.listButton {
		border: 0px;
		background-color: transparent;
	}

	input {
		display: inline-flex;
		position: relative;
		transform: translateY(-7px);
	}

	.addButton {
		height: 29px;
		display: inline-flex;
		flex-shrink: 2;
		justify-content: center;
		align-items: center;
		border-color: transparent;
		position: relative;
		/* transform: translateY(7px); fuck this and fuck css */
	}

	svg {
		vertical-align: middle;
		display: inline-block;
		line-height: normal;
	}

	.addButtonContainer {
		display: inline-flex;
		justify-content: center;
	}

	h4 {
		font-weight: 200;
		color: rgb(224, 125, 107);
	}
</style>
