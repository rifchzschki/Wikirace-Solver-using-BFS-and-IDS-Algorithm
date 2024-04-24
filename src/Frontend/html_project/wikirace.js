const availableKeywords = [
    'HTML',
    'CSS',
    'Easy Tutorials',
    'Web design tutorials',
    'JavaScripts',
    'Where to learn coding online',
    'where to learn web design',
    'how to create a website',
];

const resultsSearch = document.querySelector(".result-search");
const inputSearch1 = document.getElementById("search-bar-1");
const inputSearch2 = document.getElementById("search-bar-2");

inputSearch1.onkeyup = function() {
    let result = [];
    let input = inputSearch1.value;
    if(input.length) {
        result = availableKeywords.filter((keyword)=>{
            return keyword.toLowerCase().includes(input.toLowerCase());
        });
        console.log(result)
    }
    display1(result);

    if(!result.length) {
        resultsSearch.innerHTML = '';
    }
}

inputSearch2.onkeyup = function() {
    let result = [];
    let input = inputSearch2.value;
    if(input.length) {
        result = availableKeywords.filter((keyword)=>{
            return keyword.toLowerCase().includes(input.toLowerCase());
        });
        console.log(result)
    }
    display2(result);

    if(!result.length) {
        resultsSearch.innerHTML = '';
    }
}

function display1(result) {
    const content = result.map((list)=>{
        return "<li onclick=selectInput1(this)>" + list + "</li>";
    });

    resultsSearch.innerHTML = "<ul>" + content.join('') + "</ul>";
}

function display2(result) {
    const content = result.map((list)=>{
        return "<li onclick=selectInput2(this)>" + list + "</li>";
    });

    resultsSearch.innerHTML = "<ul>" + content.join('') + "</ul>";
}

function selectInput1(list) {
    inputSearch1.value = list.innerHTML;
    resultsSearch.innerHTML = '';
}

function selectInput2(list) {
    inputSearch2.value = list.innerHTML;
    resultsSearch.innerHTML = '';
}

function search() {
    const startPageInput = document.getElementById("search-bar-1").value;
    const targetPageInput = document.getElementById("search-bar-2").value;
    const switchButton = document.getElementById("switch");
    const searchTime = Math.floor(Math.random() * 10); 
}

const searchResult = Math.random() < 0.5; // Contoh hasil pencarian true/false
const idsPath = searchResult ? "pathIDS" : "";
const bfsPath = searchResult ? "pathBFS" : "";

const idsPathElement = document.getElementById("ids-path");
const bfsPathElement = document.getElementById("bfs-path");

if (searchResult) {
    if (switchButton.checked) {
        idsPathElement.textContent = `Found ${idsPath} with IDS of separation from ${startPageInput} to ${targetPageInput} in ${searchTime} seconds!`;
        bfsPathElement.textContent = "";
    } else {
        bfsPathElement.textContent = `Found ${bfsPath} with BFS of separation from ${startPageInput} to ${targetPageInput} in ${searchTime} seconds!`;
        idsPathElement.textContent = "";
    }
} else {
    idsPathElement.textContent = "";
    bfsPathElement.textContent = "";
}

document.getElementById("search-time").textContent = `Search time: ${searchTime} seconds`;


// Fungsi untuk menampilkan hasil pencarian di kotak hasil
function displayResults(result) {
const resultsSearch = document.querySelector(".result-search");
const content = result.map((list) => {
    return "<li>" + list + "</li>";
});

resultsSearch.innerHTML = "<ul>" + content.join("") + "</ul>";
}

// Event listener untuk tombol GO
document.querySelector(".go-button").addEventListener("click", function (event) {
event.preventDefault(); // Untuk mencegah reload halaman
search();
});

// Event listener untuk input search 1 dan 2
inputSearch1.addEventListener("keyup", function () {
let result = [];
let input = inputSearch1.value;
if (input.length) {
    result = availableKeywords.filter((keyword) => {
        return keyword.toLowerCase().includes(input.toLowerCase());
    });
}
displayResults(result);
});

inputSearch2.addEventListener("keyup", function () {
let result = [];
let input = inputSearch2.value;
if (input.length) {
    result = availableKeywords.filter((keyword) => {
        return keyword.toLowerCase().includes(input.toLowerCase());
    });
}
displayResults(result);
});