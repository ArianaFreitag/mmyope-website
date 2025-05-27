const url = "https://mmyope-app-th9iq.ondigitalocean.app";

window.onload = async function load_page() {
  let book_id = new URLSearchParams(window.location.search).get("bookId");

  let test = (document.getElementById("pub_specs").innerHTML =
    '<p>printed in black and white at <a href="#">Imprimerie Française</a> on Heavyweight Cold Press paper. 250 pages in a landscape A4 format.</p><br><p>published TEST, 2025</p><p>lifetime edition of 50</p>"');

  if (book_id == null) {
    console.log("no url params");
    book = await get_all_books();
    book_id = book.id;
  }

  sucess = await get_book(book_id);

  // console.log("2", href);
};

function replace_img(img) {
  let img_src = img.src;

  let main_img = document.getElementById("bookpage_img_0");
  main_img.src = img_src;
}

async function get_all_books() {
  try {
    const response = await fetch(url + "/getBooks");
    const res_json = await response.json();
    if (!response.ok) {
      // need to set up real statuses
      throw new Error(`Request failed`);
    }

    var newest_book = res_json.at(-1);
    var href = new URL(window.location.href);

    href.searchParams.set("bookId", newest_book.id);
    window.history.pushState("", "", href);
    console.log(res_json);
    return newest_book;
  } catch (error) {
    console.error(error.message);
  }
}

async function get_book(book_id) {
  console.log("fetching book " + book_id);
  try {
    const response = await fetch(url + "/getBook?id=" + book_id);
    const res_json = await response.json();
    if (!response.ok) {
      // need to set up real statuses
      throw new Error(`Request failed`);
    }
    console.log(res_json);
    update_content(book_id, res_json);
    return true;
  } catch (error) {
    console.error(error.message);
  }
}

function update_content(book_id, response) {
  // titles
  document.getElementById("book_title").innerText = response["title"];
  document.getElementById("book_num").innerText = response["book_num"];
  // images
  let fixed_paths = response["img_paths"]
    .replace("{", "[")
    .replace("}", "]")
    .replace(/'/g, '"');
  let i = 0;
  let element_id = "";
  for (const path of JSON.parse(fixed_paths)) {
    // replace main image with first image
    if (i == 0) {
      element_id = "bookpage_img_" + i;
      const img = document.getElementById(element_id);
      console.log(element_id);
      img.src = "images/books/" + path;
    }
    i = i + 1;
    element_id = "bookpage_img_" + i;
    const img = document.getElementById(element_id);
    img.src = "images/books/" + path;
  }
}

// async function get_file(path) {
//   console.log(path);
//   const response = await fetch(url + "/getImages?path=" + path);
//   console.log(response);
//   let blob = await response.blob();
//   const imageURL = URL.createObjectURL(blob);
//   element_id = path.split("/").at(-1).split(".")[0];
//   const img = document.getElementById(element_id);
//   img.src = imageURL;
// }
//
