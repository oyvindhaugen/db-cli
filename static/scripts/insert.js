function insertRow() {
    // amount = parseInt(amount)
    let item = document.getElementById('item')
    let amount = document.getElementById('amount')
    let data = {
        Item: item.value ,
        Amount: parseInt(amount.value)
    }
    fetch('/insert_row', {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: 'POST',
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data)
            console.log(result)
        })
    }).catch((error) => {
        console.log(error)
    })

    console.log(typeof(item.value), amount.value)
    window.location.href="/"
}