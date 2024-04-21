(async () => {
    const result = await fetch("http://localhost:8090/users?userId=2", {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        },
        //body: JSON.stringify({id: 3, name: "Dima"})
    })

    console.log(await result.json())
})()