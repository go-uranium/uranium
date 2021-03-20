// Nav bar Dropdown
document.addEventListener('DOMContentLoaded', () => {
    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);
    if ($navbarBurgers.length > 0) {
        $navbarBurgers.forEach(el => {
            el.addEventListener('click', () => {
                const target = el.dataset.target;
                const $target = document.getElementById(target);
                el.classList.toggle('is-active');
                $target.classList.toggle('is-active');
            });
        });
    }
});

// Regular Dropdown
const dropdownList = document.querySelectorAll('.dropdown:not(.is-hoverable)');
if (dropdownList.length > 0) {
    dropdownList.forEach(function (el) {
        el.addEventListener('click', function (e) {
            e.stopPropagation();
            el.classList.toggle('is-active');
        });
    });
    document.addEventListener('click', function (e) {
        dropdownList.forEach(function (el) {
            el.classList.remove('is-active');
        });
    });
}

// Alert
function alertInfo(msg) {
    bulmaToast.toast({
        message: msg,
        type: 'is-info',
        position: 'top-center',
        dismissible: true,
        duration: 4000,
        pauseOnHover: true,
        animate: {in: 'fadeIn', out: 'fadeOut'},
    })
}

function alertDanger(msg) {
    bulmaToast.toast({
        message: msg,
        type: 'is-danger',
        position: 'top-center',
        dismissible: true,
        duration: 4000,
        pauseOnHover: true,
        animate: {in: 'fadeIn', out: 'fadeOut'},
    })
}

// Common functions
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

function loggedIn() {
    return !(getCookie('token') === undefined ||
        getCookie('token').length === 0);
}


// vote for post
function updateVote(pos, neg) {
    document.getElementById("pos").innerText = pos;
    document.getElementById("neg").innerText = neg
}

function disableVote() {
    document.getElementById("pos").setAttribute("disabled", "true");
    document.getElementById("neg").setAttribute("disabled", "true")
}

function enableVote() {
    document.getElementById("pos").removeAttribute("disabled");
    document.getElementById("neg").removeAttribute("disabled")
}

function checkVote() {
    return !(document.getElementById("pos").hasAttribute("disabled") ||
        document.getElementById("neg").hasAttribute("disabled"))
}

function vote(pid, type) {
    if (loggedIn()) {
        if (checkVote()) {
            disableVote();
            const xhr = new XMLHttpRequest();
            xhr.open("POST", '/vote/post', true);
            xhr.setRequestHeader('Content-Type', 'application/json');
            xhr.onreadystatechange = function () {
                if (xhr.readyState === 4) {
                    if (xhr.status === 200) {
                        const json = JSON.parse(xhr.responseText);
                        console.log(json);
                        if (json.success) {
                            updateVote(json.pos, json.neg)
                        } else {
                            alertDanger(json.err)
                        }
                    } else {
                        alertDanger("an error occurred")
                    }
                    enableVote()
                }
            };
            xhr.send(JSON.stringify({
                type: type,
                id: parseInt(pid)
            }));
        } else {
            alertInfo("Processing... Please wait...")
        }
    } else {
        alertInfo("Please login first.")
    }
}