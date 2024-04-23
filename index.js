(async () => {
    const result = await fetch("http://localhost:4545/users/new", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify([{id: 23, name: "Example"}])
    })

    console.log(await result.json())
})()