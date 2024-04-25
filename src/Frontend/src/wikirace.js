const inputSearch1 = document.getElementById("search-bar-1");
const inputSearch2 = document.getElementById("search-bar-2");

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
          listItem.setAttribute("onclick", "selectInput1(this)");
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
          listItem.setAttribute("onclick", "selectInput2(this)");
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

function selectInput1(list) {
  inputSearch1.value = list.innerHTML;
  autocompleteList1.innerHTML = "";
}

function selectInput2(list) {
  inputSearch2.value = list.innerHTML;
  autocompleteList2.innerHTML = "";
}

function swapValues() {
  if (inputSearch1 && inputSearch2) {
    const temp = inputSearch1.value;
    inputSearch1.value = inputSearch2.value;
    inputSearch2.value = temp;

    console.log("Nilai telah ditukar:", startValue, targetValue); // Pemeriksaan apakah nilai telah ditukar
  } else {
    console.log("Salah satu atau kedua input tidak ditemukan.");
  }
}
