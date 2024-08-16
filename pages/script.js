document
  .getElementById("calcForm")
  .addEventListener("submit", async function (event) {
    event.preventDefault();

    // Get input values
    const num1 = document.getElementById("num1").value;
    const num2 = document.getElementById("num2").value;
    const operation = document.getElementById("operation").value;

    // Prepare request data
    const requestData = {
      num1: parseFloat(num1),
      num2: parseFloat(num2),
      operation: operation,
    };

    try {
      // Send request to the server
      const response = await fetch("/mats", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      });

      // Parse response
      const result = await response.json();

      // Display result
      document.getElementById("result").textContent =
        `Result: ${result.result}`;
    } catch (error) {
      console.error("Error:", error);
      document.getElementById("result").textContent = "An error occurred.";
    }
  });
