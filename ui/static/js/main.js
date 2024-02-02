let select = document.getElementById("Goods");

if(select ){
	select.addEventListener("change", myFunction);
}


function myFunction() {
	let select = document.getElementById("Goods");
	let textInput = document.getElementById("text");
	let textLabel = document.getElementById("ltext");
	if (select.value !== "") {
		textInput.style = "display: content";
		textLabel.style = "display: content";
	}
}
console.log("test")