document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".clickable-image").forEach((btn) => {
        btn.addEventListener("click", function () {

            const heroName = this.dataset.name?.trim(); 

            if (!heroName) {
                console.error("Hero name is empty!");
                return;
            }

            fetch("http://localhost:8080/heroName", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ name: heroName })
            })
            .then(response => response.json())
            .then(data => {
                console.log("Hero set successfully:", data);

                window.location.href = "http://localhost:8080/hero";
            })
            .catch(error => console.error("Error:", error));
        });
    });
});
