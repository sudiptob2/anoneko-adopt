<script lang="ts">
	export let data;
	console.log(data);
	let showFullDescription = false;
	function toggleDescription() {
		showFullDescription = !showFullDescription;
	}
</script>

<div class="grid grid-cols-1 text-center mt-3 mb-6 font-light">
	<h1 class="text-4xl">Welcome to our adoption center!</h1>
	<p>
		<span class=" font-semibold">Turning Strays into Family:</span> Every Pet is a Priority, Not an Option.
	</p>
</div>

<div class="mx-10 sm:mx-20 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-3 gap-4">
	{#each data.contents as pet}
		<div class="card glass">
			<figure class="h-[150px]">
				<img class="h-[150px] object-cover" src={pet.thumbImageUri} alt="pet!" />
			</figure>
			<div class="card-body">
				<div>
					<h2 class="card-title">{pet.name}</h2>
					<div class="badge badge-neutral">{pet.lawchaos}</div>
					<div class="badge badge-accent">{pet.goodevil}</div>
					<div class="badge badge-info">age: {pet.age}</div>
				</div>

				<div class="text-justify">
					<p>
						{pet.description.substring(0, 150)}...
					</p>
				</div>

				<div class="card-actions justify-center">
					<button class="btn btn-primary">Adopt</button>
					<button class="btn btn-secondary ml-2">View Profile</button>
				</div>
			</div>
		</div>
	{/each}
</div>
