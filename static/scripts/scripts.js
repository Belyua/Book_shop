$(document).ready(function() {
    // Fetch and display books on index.html load
    $.get("/find-books", function(data) {
        var bookList = $("#book-list");
        bookList.empty();
        data.Books.forEach(function(book) {
            bookList.append("<li>" + book.Title + " by " + book.Author + "</li>");
        });
    });

    // Handle Create Book form submission
    $("#create-form").submit(function(event) {
        event.preventDefault();
        var formData = {
            title: $("#title").val(),
            author: $("#author").val()
        };
        $.post("/create-book", formData, function() {
            window.location.href = "/";
        });
    });

    // Handle Update Book form submission
    $("#update-form").submit(function(event) {
        event.preventDefault();
        var formData = {
            title: $("#title").val(),
            author: $("#author").val()
        };
        var bookId = $("#update-id").val();
        $.ajax({
            url: "/update-book/" + bookId,
            method: "PUT",
            data: JSON.stringify(formData),
            contentType: "application/json",
            success: function() {
                window.location.href = "/";
            }
        });
    });

    // Fetch book details and populate Update form
    $("#book-list").on("click", "li", function() {
        var bookTitle = $(this).text().split(" by ")[0];
        $.get("/find-book-by-title/" + bookTitle, function(data) {
            $("#update-id").val(data.data.id);
            $("#title").val(data.data.title);
            $("#author").val(data.data.author);
        });
    });
});
