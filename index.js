(async () => {
    const result = await fetch("http://localhost:4545/users/update", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            id: 23,
            user: {
                id: 4,
                name: "Updated"
            }
        })
    })

    console.log(await result.json())
})()