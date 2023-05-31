// function signUp() {
//   var xhr = new XMLHttpRequest();

//   xhr.open('POST', '/api/signup');
//   xhr.setRequestHeader('Content-Type', 'application/json');
//   xhr.onload = function() {
//     if (xhr.status === 200) {
//       var res = JSON.parse(xhr.responseText);
//       if (res.success == "true") {
//         // document.getElementById("userinfo").innerHTML = "Hello, " + res.username       
//       } else {
//         // document.getElementById("userinfo").innerHTML = ""
//       }
//     }
//   };

//   xhr.send(JSON.stringify({
//         username: document.getElementById("username").value,
//         password: document.getElementById("password").value
//   }));
// }

function showMenu() {
        document.getElementById("menu").style.display = "unset";
}

function hideMenu() {
        document.getElementById("menu").style.display = "none";
}

function jumpTo(eid) {
    var jump = document.getElementById(eid);
jump.scrollIntoView({
            behavior: 'auto',
            block: 'center',
            inline: 'center'
});
    hideMenu();
}


