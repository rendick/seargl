document.addEventListener("DOMContentLoaded", function () {
    var xmpElement = document.getElementById("output");
    if (xmpElement) {
      var content = xmpElement.innerHTML;
      var updatedContent = content
        .replace(/\[1m/g, "")
        .replace(/\[0m/g, "");
      xmpElement.innerHTML = updatedContent;
    }
  });