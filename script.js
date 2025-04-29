fetch("puzzles.json")
.then(response => response.json())
.then(data => {
    const table = document.getElementById("problems-table-table")
    data.forEach(puzzle =>{
        const row = document.createElement("tr")
        row.innerHTML = `
        <td>${puzzle.name}</td>
        <td><a href="${puzzle.python}">link</a></td>
        <td><a href="${puzzle.golang}">link</td>
        `
        table.appendChild(row)
    })
})