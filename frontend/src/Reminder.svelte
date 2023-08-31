<script>
	import autoAnimate from "@formkit/auto-animate";

	let reminderList = [];
	let newReminderText = "";
	let emptyReminderWarning = "";

	// Add a new reminder to the list
	function addReminder() {
		if (newReminderText == "") {
			emptyReminderWarning = "Please enter a reminder";
			return;
		} else {
			emptyReminderWarning = "";
		}

		const id = reminderList.length + 1;
		const reminder = { id, text: newReminderText, completed: false };
		reminderList = [...reminderList, reminder];
		newReminderText = "";
	}

	// Delete a reminder from the list
	function deleteReminder(id) {
		reminderList = reminderList.filter((reminder) => reminder.id !== id);
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
