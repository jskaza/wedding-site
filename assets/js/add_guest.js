let i = 2;
function addGuest() {

    var html = `

<div class="card w-11/12 bg-base-100 shadow-xl">
    <div class="card-body">
        <div class="card-actions justify-end">
            <button type="button" class="btn btn-square btn-sm">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M6 18L18 6M6 6l12 12" />
                </svg>
            </button>
        </div>
        <input required name="people[${i}][name]" form="rsvp" type="text" placeholder="Guest Name" class="
                mt-0
                block
                w-full
                px-0.5
                border-0 border-b-2 border-gray-200
                focus:ring-0 focus:border-black
              " />
        <select required form="rsvp" name="people[${i}][meal]" class=" block
                w-full
                mt-0
                px-0.5
                border-0 border-b-2 border-gray-200
                focus:ring-0 focus:border-black">
            <option>Chicken</option>
            <option>Steak</option>
        </select>
    </div>
</div>
</div>
`;


    var guests = document.getElementById("cards");
    var addButton = document.getElementById("addButton");
    var card = document.createElement("div");
    card.innerHTML = html;

    guests.insertBefore(card, addButton);
    window.scrollTo(0, document.body.scrollHeight, 'smooth');
    i++;
}
