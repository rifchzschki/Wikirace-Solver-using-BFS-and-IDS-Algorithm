const availableKeywords = [
  "HTML",
  "CSS",
  "Easy Tutorials",
  "Web design tutorials",
  "JavaScripts",
  "Where to learn coding online",
  "where to learn web design",
  "how to create a website",
];

const resultsSearch = document.querySelector(".result-search");
const inputSearch1 = document.getElementById("search-bar-1");
const inputSearch2 = document.getElementById("search-bar-2");

// const searchInput = document.getElementById("search");
const autocompleteList1 = document.getElementById("autocomplete-list-1");
const autocompleteList2 = document.getElementById("autocomplete-list-2");

function autoCompleteSearchBar1() {
  const query = this.value;
  if (query.length === 0) {
    autocompleteList1.innerHTML = ""; // Menghapus semua elemen li jika query kosong
    return;
  }
  fetch(`/autocomplete?query=${query}`)
    .then((response) => response.json())
    .then((data) => {
      const currentQuery = this.value;
      if (query !== currentQuery) {
        // Jika query saat ini sudah berubah, hentikan eksekusi
        return;
      }
      autocompleteList1.innerHTML = ""; // Menghapus semua elemen li
      if (data.length === 0) {
        // Jika data kosong, tambahkan pesan informasi ke autocompleteList1
        const messageItem = document.createElement("li");
        messageItem.textContent = "No results found";
        autocompleteList1.appendChild(messageItem);
      } else {
        // Jika data tidak kosong, tambahkan item-item dari data ke autocompleteList1
        data.forEach((item) => {
          const listItem = document.createElement("li");
          listItem.textContent = item;
          autocompleteList1.appendChild(listItem);
        });
      }
    })
    .catch((error) =>
      console.error("Error fetching autocomplete data:", error)
    );
}
function autoCompleteSearchBar2() {
  const query = this.value;
  if (query.length === 0) {
    autocompleteList2.innerHTML = ""; // Menghapus semua elemen li jika query kosong
    return;
  }
  fetch(`/autocomplete?query=${query}`)
    .then((response) => response.json())
    .then((data) => {
      const currentQuery = this.value;
      if (query !== currentQuery) {
        // Jika query saat ini sudah berubah, hentikan eksekusi
        return;
      }
      autocompleteList2.innerHTML = ""; // Menghapus semua elemen li
      if (data.length === 0) {
        // Jika data kosong, tambahkan pesan informasi ke autocompleteList2
        const messageItem = document.createElement("li");
        messageItem.textContent = "No results found";
        autocompleteList2.appendChild(messageItem);
      } else {
        // Jika data tidak kosong, tambahkan item-item dari data ke autocompleteList2
        data.forEach((item) => {
          const listItem = document.createElement("li");
          listItem.textContent = item;
          autocompleteList2.appendChild(listItem);
        });
      }
    })
    .catch((error) =>
      console.error("Error fetching autocomplete data:", error)
    );
}

inputSearch1.addEventListener("input", autoCompleteSearchBar1);
inputSearch2.addEventListener("input", autoCompleteSearchBar2);

function swapValues() {
  document.getElementById("search-bar-1").value = inputSearch2;
  document.getElementById("search-bar-2").value = inputSearch1;
  if (inputSearch1 && inputSearch2) {
    // Memeriksa apakah kedua input ditemukan
    var startValue = inputSearch1.value;
    var targetValue = inputSearch2.value;

    // Menukar nilai
    inputSearch1.value = targetValue;
    inputSearch2.value = startValue;

    console.log("Nilai telah ditukar:", startValue, targetValue); // Pemeriksaan apakah nilai telah ditukar
  } else {
    console.log("Salah satu atau kedua input tidak ditemukan.");
  }
}

inputSearch1.onkeyup = function () {
  let result = [];
  let input = inputSearch1.value;
  if (input.length) {
    result = availableKeywords.filter((keyword) => {
      return keyword.toLowerCase().includes(input.toLowerCase());
    });
    console.log(result);
  }
  display1(result);

  if (!result.length) {
    resultsSearch.innerHTML = "";
  }
};

inputSearch2.onkeyup = function () {
  let result = [];
  let input = inputSearch2.value;
  if (input.length) {
    result = availableKeywords.filter((keyword) => {
      return keyword.toLowerCase().includes(input.toLowerCase());
    });
    console.log(result);
  }
  display2(result);

  if (!result.length) {
    resultsSearch.innerHTML = "";
  }
};

function display1(result) {
  const content = result.map((list) => {
    return "<li onclick=selectInput1(this)>" + list + "</li>";
  });

  resultsSearch.innerHTML = "<ul>" + content.join("") + "</ul>";
}

function display2(result) {
  const content = result.map((list) => {
    return "<li onclick=selectInput2(this)>" + list + "</li>";
  });

  resultsSearch.innerHTML = "<ul>" + content.join("") + "</ul>";
}

function selectInput1(list) {
  inputSearch1.value = list.innerHTML;
  resultsSearch.innerHTML = "";
}

function selectInput2(list) {
  inputSearch2.value = list.innerHTML;
  resultsSearch.innerHTML = "";
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

document.getElementById(
  "search-time"
).textContent = `Search time: ${searchTime} seconds`;

// Fungsi untuk menampilkan hasil pencarian di kotak hasil
function displayResults(result) {
  const resultsSearch = document.querySelector(".result-search");
  const content = result.map((list) => {
    return "<li>" + list + "</li>";
  });

  resultsSearch.innerHTML = "<ul>" + content.join("") + "</ul>";
}

// Event listener untuk tombol GO
document
  .querySelector(".go-button")
  .addEventListener("click", function (event) {
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
