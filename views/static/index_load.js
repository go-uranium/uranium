window.onscroll = function () {
    var scrollHeight = document.documentElement.scrollHeight;
    var scrollTop = document.documentElement.scrollTop;
    var clientHeight = document.documentElement.clientHeight;

    if (clientHeight + scrollTop >= scrollHeight) {
        console.log("===加载更多内容……===");
    }
};