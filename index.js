(async () => {
    const result = await fetch("http://localhost:4545/users/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify([{id: 1, name: "Snowik"}])
    })

    console.log(await result.json())
})()