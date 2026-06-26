# PSEUDOCODE

START PROGRAM

LOAD HTML template from "templates/index.html"

DEFINE structure Data
    Result : string
    Error  : string

MAIN
    Serve static files from "/static/"
    Register "home" function for route "/"

    Print "server running on http://localhost:8001"

    Start HTTP server on port 8001
END MAIN


FUNCTION home(responseWriter, request)

    Parse form data
    IF parsing fails THEN
        Return HTTP 400 error
    END IF

    IF request method is GET THEN
        Render template with no data
        RETURN
    END IF

    IF request method is POST THEN

        Read form values:
            a
            b
            c

        Convert a, b, c from strings to floating-point numbers

        Call complexxQuadraticEquation(a, b, c)

        IF error occurs THEN
            Return HTTP 400 error
            RETURN
        END IF

        Create result string:
            "Roots: root1, root2"

        Create Data object:
            Result = result string

        Render template with Data object

    ELSE
        Return HTTP 405 (Method Not Allowed)
    END IF

END FUNCTION


FUNCTION complexxQuadraticEquation(a, b, c)

    IF a = 0 THEN
        Return error:
            "coefficient 'a' cannot be zero"
    END IF

    Convert a, b, c to complex numbers:
        A, B, C

    Calculate discriminant:
        discriminant = B² - 4AC

    Calculate square root of discriminant:
        sqrtDiscriminant = √discriminant

    Calculate first root:
        root1 = (-B + sqrtDiscriminant) / (2A)

    Calculate second root:
        root2 = (-B - sqrtDiscriminant) / (2A)

    Return root1, root2, no error

END FUNCTION


PROGRAM FLOW

1. User opens browser at localhost:8001.
2. GET request is sent.
3. Server displays HTML form.
4. User enters coefficients a, b, and c.
5. Form sends POST request.
6. Server reads coefficients.
7. Server computes quadratic roots using:
       x = (-b ± √(b² - 4ac)) / 2a
8. Roots are returned (real or complex).
9. Result is displayed on the webpage.

END PROGRAM

# IMPLEMENTATION

IMPLEMENTATION OF A WEB-BASED QUADRATIC EQUATION SOLVER

1. Initialize the application by loading the HTML template file (`index.html`).

2. Define a data structure containing:

   * Result: stores the calculated roots.
   * Error: stores any error message.

3. Start the HTTP server:

   * Configure static file serving for CSS, JavaScript, and images.
   * Register the home route (`/`).
   * Listen for incoming requests on port 8001.

4. When a request arrives at the home route:

   a. Parse the submitted form data.

   b. If form parsing fails:

   * Return a "Bad Request" error.
   * Stop processing.

5. Handle GET requests:

   * Render and display the HTML form.
   * Wait for user input.

6. Handle POST requests:

   a. Retrieve the values of coefficients:

   * a
   * b
   * c

   b. Convert the coefficient values from strings to floating-point numbers.

   c. Call the quadratic equation solver function using:

   * aFloat
   * bFloat
   * cFloat

   d. If the solver returns an error:

   * Display the error message.
   * Stop processing.

   e. If successful:

   * Format the roots into a result string.
   * Store the result in the data structure.
   * Render the template with the calculated result.

7. Reject unsupported HTTP methods by returning a "Method Not Allowed" error.

---

QUADRATIC EQUATION SOLVER IMPLEMENTATION

Input:
a, b, c

Output:
Two roots (real or complex)

Steps:

1. Check whether coefficient a equals zero.

   * If true:
     Return an error because the equation is not quadratic.

2. Convert coefficients to complex-number format.

3. Compute the discriminant:
   D = b² − 4ac

4. Compute the square root of the discriminant:
   √D

5. Compute the first root:
   x₁ = (-b + √D) / (2a)

6. Compute the second root:
   x₂ = (-b - √D) / (2a)

7. Return both roots.

---

SYSTEM WORKFLOW

User → Opens webpage

```
  ↓
```

Server displays form

```
  ↓
```

User enters values for a, b, and c

```
  ↓
```

Form submits POST request

```
  ↓
```

Server validates input

```
  ↓
```

Quadratic solver computes roots

```
  ↓
```

Results returned to webpage

```
  ↓
```

User views the calculated roots
