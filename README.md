## Scrape-Us : Scrape us daddy

TODOS

1. **URL Fetching**: This is the first step in web scraping where you fetch the URL of the webpage you want to scrape.

2. **HTML Parsing**: After fetching the webpage, the next step is to parse the HTML of the webpage.

   - Do not Use a library like `goquery` to parse the HTML.

3. **Data Extraction**: Once the HTML is parsed, you can then extract the data you need.

   - Identify the HTML elements that contain the data you want to scrape and extract the data.

4. **Data Transformation**: After extracting the data, transform it into a format that's useful for your purposes.

   - Implement data transformation functions.

5. **Data Storage**: Once the data is transformed, you can then store it.

   - Decide on a storage method (e.g., database, file system) and implement it.

6. **Error Handling**: handle errors gracefully

   - Implement error handling throughout your code.

7. **Concurrency**:

   - Use go-routines and channels to implement concurrency.

8. **Rate Limiting**: To avoid getting blocked by the website you're scraping, you might need to implement rate limiting.

   - Implement rate limiting.

9. **Logging**:

   - Implement logging.

10. **Configuration**:
    Task: Implement configuration options.

11. **Testing**:
    - Task: Write tests
