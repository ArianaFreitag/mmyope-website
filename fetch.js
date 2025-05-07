const url = "http://localhost:3333";

window.onload = async function load_page() {
  // await get_file();
  book = await get_all_books();
  sucess = await get_book(book.id);
};

async function get_file(path) {
  console.log(path);
  const response = await fetch(url + "/getImages?path=" + path);
  console.log(response);
  let blob = await response.blob();
  const imageURL = URL.createObjectURL(blob);
  element_id = path.split("/").at(-1).split(".")[0];
  const img = document.getElementById(element_id);
  img.src = imageURL;
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
    // convert to real string
    let fixed_paths = res_json["img_paths"]
      .replace("{", "[")
      .replace("}", "]")
      .replace(/'/g, '"');
    // console.log(JSON.parse(fixed_paths));

    for (const path of JSON.parse(fixed_paths)) {
      // get_file(path);
      // const imageURL = URL.createObjectURL(blob);
      element_id = path.split("/").at(-1).split(".")[0];
      const img = document.getElementById(element_id);
      img.src = "images/" + path;
    }

    return true;
  } catch (error) {
    console.error(error.message);
  }
}
