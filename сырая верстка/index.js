
var checkbox = document.querySelector(".show-completed input[type=checkbox]");
checkbox.addEventListener('change', function() {
  var tasks = document.getElementsByClassName("task-item");  
  if (this.checked) {
    for (task in tasks){
        if (document.querySelector("input").checked){
            document.getElementById("changeview").innerHTML = ``
        }
    }
  } else {
    for (task in tasks){
        if (document.querySelector("input").checked){
            document.getElementById("changeview").innerHTML = `
            .task-item {
                display: none;
            }`
        }
    }
  }
});