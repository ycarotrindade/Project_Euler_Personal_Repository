function filter_data(e){
    clear_table()
    if (e == null || e.value == ""){
        fetch("puzzles.json")
        .then(response => response.json())
        .then(data => {
            const table = document.getElementById("problems-table-table")
            data.forEach(puzzle =>{
                const row = document.createElement("tr")
                row.innerHTML = `
                <td>${puzzle.name}</td>
                <td><a href="${puzzle.python}" target="_blank">link</a></td>
                <td><a href="${puzzle.golang}" target="_blank">link</td>
                `
                table.appendChild(row)
            })
        })
    }else{
        fetch("puzzles.json")
        .then(response => response.json())
        .then(data => {
            const table = document.getElementById("problems-table-table")
            data.forEach(puzzle =>{
                if (puzzle.name.includes(e.value)){
                    const row = document.createElement("tr")
                    row.innerHTML = `
                    <td>${puzzle.name}</td>
                    <td><a href="${puzzle.python}">link</a></td>
                    <td><a href="${puzzle.golang}">link</td>
                    `
                    table.appendChild(row)
                }
            })
        })
    }
}

function clear_table(){
    const table = document.getElementById("problems-table-table")
    table.innerHTML = `
    <tr>
        <th>Problem Name</th>
        <th>Python</th>
        <th>Golang</th>
    </tr>
    `
}