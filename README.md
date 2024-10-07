<h1>Go Temperature Converter</h1>

<p>This project is a simple temperature converter written in Go (Golang). It allows users to convert temperatures between Celsius and Fahrenheit via the command line. The program was built as a learning exercise to discover Go programming.</p>

<h2>What I Learned</h2>

<ul>
    <li><strong>Basic Go Syntax:</strong> I learned how to write and structure a simple Go program, including the use of <code>package main</code>, <code>import</code>, and functions.</li>
    <li><strong>Using Loops:</strong> I implemented loops to handle user input, ensuring that invalid entries (such as anything other than "C" or "F") were managed and re-prompted until correct.</li>
    <li><strong>Conditional Statements:</strong> I used <code>if</code> and <code>else</code> statements to process the logic for temperature conversion based on user input.</li>
    <li><strong>Basic I/O:</strong> I practiced how to interact with users via the command line using <code>fmt.Scanln</code> for capturing input and <code>fmt.Printf</code> for formatting and displaying output.</li>
    <li><strong>Data Validation:</strong> I learned how to validate user input (for example, ensuring only "C" or "F" is accepted for temperature types).</li>
    <li><strong>String Manipulation:</strong> I worked with strings in Go, especially using functions from the <code>strings</code> package, such as <code>ToUpper</code>, to standardize input.</li>
    <li><strong>Function Creation and Usage:</strong> I created multiple functions to break down the code into reusable parts, improving code organization and readability.</li>
    <li><strong>Loop Control:</strong> I created a continuous loop to ask users if they want to perform another conversion, utilizing <code>for</code> loops and <code>break</code> statements.</li>
</ul>

<h2>How to Run the Program</h2>

<ol>
    <li>Make sure Go is installed on your machine. You can download it from <a href="https://golang.org/dl/">https://golang.org/dl/</a>.</li>
    <li>Clone this repository:
        <pre><code>git clone https://github.com/yourusername/repository-name.git</code></pre>
    </li>
    <li>Navigate to the project directory:
        <pre><code>cd repository-name</code></pre>
    </li>
    <li>Run the program:
        <pre><code>go run main.go</code></pre>
    </li>
</ol>

<p>The program will prompt you to choose either Celsius or Fahrenheit, and then ask you for the temperature to convert. You can continue to convert temperatures until you type "NO" when asked if you want to convert another temperature.</p>

<h2>Conclusion</h2>

<p>This project provided valuable hands-on experience in Go and helped strengthen my understanding of basic programming concepts such as loops, conditionals, and user interaction in the command line.</p>
