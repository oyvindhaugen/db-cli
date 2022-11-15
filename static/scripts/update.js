//This tells the backend to update a row at given ID with given Item and Amount.
function updateRow() {
    let id = parseURL()
    id = parseInt(id)
    let item = document.getElementById('item')
    let amount = document.getElementById('amount')
    let userId = localStorage.getItem('Id')
    let data = {
        Id: id ,
        Item: item.value ,
        Amount: parseInt(amount.value) ,
        UserId: parseInt(userId)
    }
    console.log(data)
    fetch('/update_row', {
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
    setTimeout(() => {document.location.href="/"}, 350)
}
//This parses the ID stored in the URL
function parseURL() {
    let url = document.location.href, params = url.split('?')[1].split('&'), data = {}, tmp;
    for (let i = 0, l = params.length; i < l; i++) {
        tmp = params[i].split('=')
        data[tmp[0]] = tmp[1]
    }
    return tmp[1]
}