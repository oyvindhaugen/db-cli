
//This accesses the JSON file containing the data from the database
let oXHR = new XMLHttpRequest()
oXHR.onreadystatechange = reportStatus
oXHR.open("GET", "../selectQuery.json", true)
oXHR.send()
function reportStatus() {
    if (oXHR.readyState === 4) {
        createTableFromJSON(this.responseText)
    }
}
//This uses the data from the JSON file to create a dynamic table.
function createTableFromJSON(jsonData) {
    let arrShopping = []
    arrShopping = JSON.parse(jsonData)

    let col = []
    for (let i = 0; i < arrShopping.length; i++) {
        for (let key in arrShopping[i]) {
            if (col.indexOf(key) === -1) {
                col.push(key)
            }
        }
    }

    let table = document.createElement("table")
    let tr = table.insertRow(-1)

    for (let i = 0; i < col.length + 2; i++) {
        let th = document.createElement("th")
        th.innerHTML = col[i]
        tr.appendChild(th)
        if (i === 0) {
            th.classList.add("d-none")
        }
        if (i === 2) {
            th.classList.add("centerAmount")
            th.classList.add("amountHead")
        }
        if (i === 1) {
            th.classList.add("centerAmount")
            th.classList.add("itemHead")

        }
        if (i === 3) {
            th.innerHTML = 'delBtn'
            th.classList.add("d-none")
            tr.appendChild(th)
        }
        if (i === 4) {
            th.innerHTML = 'updtBtn'
            th.classList.add("d-none")
            tr.appendChild(th)
        }

    }


    for (let i = 0; i < arrShopping.length; i++) {
        tr = table.insertRow(-1)
        tr.setAttribute('id', "tr" + (i + 1))
        for (let j = 0; j < col.length + 2; j++) {
            let tabCell = tr.insertCell(-1)
            tabCell.innerHTML = arrShopping[i][col[j]]
            if (j === 0) {
                tabCell.classList.add("d-none")
            }
            if (j === 2) {
                tabCell.classList.add("centerAmount")
                tabCell.classList.add("amountRow")
            }
            if (j === 1) {
                tabCell.classList.add("centerAmount")
                tabCell.classList.add("itemRow")
            }
            //checks if it's the last row and if it is, it adds the bottom class
            if (i + 1 >= arrShopping.length) {
                tabCell.classList.add("bottom")
            }

            if (j === 3) {
                tabCell.innerHTML = " "
                let img = document.createElement('img')
                img.setAttribute('src', 'https://cdn-icons-png.flaticon.com/512/484/484662.png')
                img.setAttribute('width', '20px')
                tabCell.setAttribute('type', 'button')
                tabCell.setAttribute('value', 'Delete Row')
                tabCell.classList.add('align-self-center', 'btn', 'btn-danger')
                tabCell.setAttribute('onclick', 'deleteRow(this.id)')
                tabCell.setAttribute('id', 'd' + (arrShopping[i][col[0]]))
                tabCell.appendChild(img)
            }
            if (j === 4) {
                tabCell.innerHTML = " "
                let img = document.createElement('img')
                img.setAttribute('src', 'https://cdn-icons-png.flaticon.com/512/1250/1250925.png')
                img.setAttribute('width', '20px')
                tabCell.setAttribute('type', 'button')
                tabCell.classList.add('align-self-center', 'btn', 'btn-success')
                tabCell.setAttribute('onclick', 'rediToUpdt(this.id)')
                tabCell.setAttribute('id', 'u' + (arrShopping[i][col[0]]))
                tabCell.appendChild(img)
            }
        }
    }
    let divContainer = document.getElementById('showTable')
    divContainer.innerHTML = ""
    divContainer.appendChild(table)
    let divContainerBt = document.getElementById('showBtn')
    divContainerBt.innerHTML = ""
}
//This tells the backend to delete a row at given ID
function deleteRow(id) {
    if (window.confirm('Are you sure you want to delete this item?')) {
        id = trim(id)
        id = parseInt(id)
        let data = {
            Id: id
        }
        fetch("/delete_row", {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            method: "POST",
            body: JSON.stringify(data)
        }).then((response) => {
            response.text().then(function (data) {
                let result = JSON.parse(data)
                console.log(result)
            })
        }).catch((error) => {
            console.log(error)
        })
        setTimeout(() => { window.location.reload() }, 350)
    }
}
//This trims the first char in given string, used for trimming ID
function trim(s) {
    let sTrimmed = s.substring(1)
    return sTrimmed
}
//This redirects to update.html and also stores the given ID in the URL
function rediToUpdt(id) {
    id = trim(id)
    url = 'http://127.0.0.1:5500/update.html?id=' + encodeURIComponent(id)
    document.location.href = url
}