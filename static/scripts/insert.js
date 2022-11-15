function insertRow() {
    let item = document.getElementById('item')
    let amount = document.getElementById('amount')
    let userId = localStorage.getItem('Id')
    let data = {
        Item: item.value,
        Amount: parseInt(amount.value),
        UserId: parseInt(userId)
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
    setTimeout(() => { document.location.href = "/" }, 350)
}
