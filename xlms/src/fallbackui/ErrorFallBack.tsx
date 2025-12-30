

function ErrorFallback({ error }) {
    console.log(error)
    return (
        <div role="alert">
            <h2>Something went wrong:</h2>
        </div>
    );
}
export default ErrorFallback