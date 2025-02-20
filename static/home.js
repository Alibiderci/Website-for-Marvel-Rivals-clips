document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".hero img").forEach(img => {
        img.addEventListener("click", function (event) {
            const heroName = event.target.alt.replace(/\s+/g, '-').toLowerCase(); // Convert "Black Panther" → "black-panther"
            window.location.href = `/hero?name=${heroName}`; // Redirect to dynamic page
        });
    });
});
