function gebi(id) {
    return document.getElementById(id);
}

async function getmenu(id) {
    uri = `/api/page/menu/${id}`;
    const resp = await apiGet(uri)
    const ul = document.createElement("ul")
    Object.entries(resp).forEach(([key, menu]) => {
        for (let i=0; i<menu.length; i++) {
            const link = document.createElement("a");
            link.href = menu[i].href;
            link.text = menu[i].title;
            link.id = menu[i].id;
            link.addEventListener('click', (e)=>{
                e.preventDefault();
                e.stopPropagation();
                getsubmenu(menu[i].submenu);
            });
            const li = document.createElement("li");
            li.appendChild(link);
            ul.appendChild(li);
        }
    });
    gebi("menu").appendChild(ul);
    return
}

async function getsubmenu(submenu) {
    const ul = document.createElement("ul")
    for (let i=0; i<submenu.length; i++) {
        const link = document.createElement("a");
        link.href = submenu[i].href;
        link.title = submenu[i].title;
        link.id = submenu[i].id;
        link.addEventListener('click', (e)=>{
            e.preventDefault();
            e.stopPropagation();
        });
        const li = document.createElement("li");
        li.appendChild(link);
        ul.appendChild(li);
    }
    gebi("submenu").appendChild(ul);
}

getmenu("main");

