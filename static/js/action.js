
const keyword = document.querySelector("#search input.keyword")
const container = document.querySelector("table#container tbody")
const pagination = document.getElementById("page")

const BASEURL = new URL("","http://localhost:8080/")


buttonAction()
// Function untuk menangani pencarian
const handleSearch = () => {
    BASEURL.pathname = "search";
    BASEURL.searchParams.set("q", keyword.value);

    fetch(BASEURL.href, {
        method: "GET"
    })
    .then((res) => {
        if (!res.ok) {
            throw new Error("Error Search Students");
        }
        return res.json();
    })
    .then((data) => {
        const students = data.Students;
        const page = data.Page;
        if (students == null || students.length == 0) {
            container.innerHTML = `NO FOUND`;
        } else {
            container.innerHTML = ``;
            students.forEach((student) => {
                container.innerHTML += `
                    <tr>
                        <td>${student.NIS}</td>
                        <td>${student.Name}</td>
                        <td>${student.Jurusan}</td>
                        <td>
                            <a href="#editEmployeeModal" class="edit" data-toggle="modal"><i class="material-icons" data-toggle="tooltip" title="Edit">&#xE254;</i></a>
                            <a href="#deleteEmployeeModal" class="delete" data-toggle="modal"><i class="material-icons" data-toggle="tooltip" title="Delete">&#xE872;</i></a>
                        </td>
                    </tr>
                `;
            });

            pagination.innerHTML = `
                <div class="hint-text">Showing <b>10</b> out of <b>${page.TotalPage}</b> entries</div>
                <ul class="pagination">
                    ${page.Current > 1 ? `<li class="page-item"><a href="?p=${page.Previous}">Previous</a></li>` : `<li class="page-item disabled"><span>Previous</span></li>`}
                    ${page.Current > 1 ? `<li class="page-item"><a href="?p=${page.Previous}" class="page-link">${page.Previous}</a></li>` : ``}
                    <li class="page-item active"><a href="?p=${page.Current}" class="page-link">${page.Current}</a></li>
                    ${page.Current < page.TotalPage ? `<li class="page-item"><a href="?p=${page.Next}" class="page-link">${page.Next}</a></li>` : ``}
                    ${page.Current < page.TotalPage ? `<li class="page-item"><a href="?p=${page.Next}" class="page-link">Next</a></li>` : `<li class="page-item disabled"><span>Next</span></li>`}
                </ul>`;

            buttonAction()
        }
    })
    .catch((error) => {
        console.error(error);
        Swal.fire({
            icon: "error",
            title: "Oops...",
            text: "Error Search Students",
        });
    });
};

// Function untuk menangani klik pada pagination
const handlePaginationClick = (event) => {
    event.preventDefault();
    if (event.target.tagName === "A") {
        const pageLink = event.target.getAttribute("href");
        if (pageLink) {
            BASEURL.pathname = "search";
            BASEURL.searchParams.set("q", keyword.value);
            BASEURL.searchParams.set("p", pageLink.split("=")[1]); // Ambil nilai parameter p dari href

            fetch(BASEURL.href, {
                method: "GET"
            })
            .then((res) => {
                if (!res.ok) {
                    throw new Error("Error Search Students");
                }
                return res.json();
            })
            .then((data) => {
                const students = data.Students;
                const page = data.Page;
                console.log(page);
                if (students == null || students.length === 0) {
                    container.innerHTML = `NO FOUND`;
                } else {
                    container.innerHTML = ``;
                    students.forEach((student) => {
                        container.innerHTML += `
                            <tr>
                                <td>${student.NIS}</td>
                                <td>${student.Name}</td>
                                <td>${student.Jurusan}</td>
                                <td>
                                    <a href="#editEmployeeModal" class="edit" data-toggle="modal"><i class="material-icons" data-toggle="tooltip" title="Edit">&#xE254;</i></a>
                                    <a href="#deleteEmployeeModal" class="delete" data-toggle="modal"><i class="material-icons" data-toggle="tooltip" title="Delete">&#xE872;</i></a>
                                </td>
                            </tr>
                        `;
                    });

                    pagination.innerHTML = `
                        <div class="hint-text">Showing <b>10</b> out of <b>${page.TotalPage}</b> entries</div>
                        <ul class="pagination">
                            ${page.Current > 1 ? `<li class="page-item"><a href="?p=${page.Previous}">Previous</a></li>` : `<li class="page-item disabled"><span>Previous</span></li>`}
                            ${page.Current > 1 ? `<li class="page-item"><a href="?p=${page.Previous}" class="page-link">${page.Previous}</a></li>` : ``}
                            <li class="page-item active"><a href="?p=${page.Current}" class="page-link">${page.Current}</a></li>
                            ${page.Current < page.TotalPage ? `<li class="page-item"><a href="?p=${page.Next}" class="page-link">${page.Next}</a></li>` : ``}
                            ${page.Current < page.TotalPage ? `<li class="page-item"><a href="?p=${page.Next}" class="page-link">Next</a></li>` : `<li class="page-item disabled"><span>Next</span></li>`}
                        </ul>`;
                        buttonAction()
                }
            })
            .catch((error) => {
                console.error(error);
                Swal.fire({
                    icon: "error",
                    title: "Oops...",
                    text: "Error Search Students",
                });
            });
        }
    }
};

// Tambahkan event listener untuk pencarian
keyword.addEventListener("keyup", handleSearch);

// Tambahkan event listener untuk klik pada pagination
pagination.addEventListener("click", handlePaginationClick);



























function buttonAction() {
    const buttons = document.querySelectorAll("a.delete")
    const buttonEdit = document.querySelectorAll("a.edit")
    const submit = document.querySelector("input.submitDelete")
    const modal = document.getElementById("editEmployeeModal")
    buttons.forEach((button,i) => {
        button.addEventListener("click",() => {
            const student = button.parentElement.parentElement,
            NIS = student.children[0].textContent
            submit.addEventListener("click",() => {
                BASEURL.pathname = `delete/${NIS}`
        
                fetch(BASEURL.href, {
                    method: "POST",
                })
                .then((res) => {
                    if (res.ok) {
                        student.remove();
                        window.location.reload()  
                    } else {
                        Swal.fire({
                            icon: "error",
                            title: "Oops...",
                            text: "Error Deleted Students",
                          });
                    }
                })
                .catch((error) => {
                    console.error('Fetch error:', error);
                    Swal.fire({
                        icon: "error",
                        title: "Oops...",
                        text: error,
                      });
                });
            })
            
        })
    })
    buttonEdit.forEach((button,i) => {
        button.addEventListener("click",() => {
            const student = button.parentElement.parentElement,
            NIS = student.children[0].textContent,
            name = student.children[1].textContent,
            jurusan = student.children[2].textContent
    
    
            const form = modal.querySelector("form")
            form.setAttribute("action",`/update/${NIS}`)
            form.setAttribute("method",`post`) 
            const input = modal.querySelector("form .modal-body")
            const inputNIS = input.querySelector(".form-group input[name='NIS']"),
             inputName = input.querySelector(".form-group input[name='Name']"),
             inputJurusan = input.querySelector(".form-group input[name='Jurusan']")
    
            inputNIS.setAttribute("value",NIS)
            inputName.setAttribute("value",name)
            inputJurusan.setAttribute("value",jurusan)
    
        })
    })
     
}

