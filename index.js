(async () => {
    const result = await fetch("http://localhost:8090/headers", {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    })

    console.log(await result.json())
})()