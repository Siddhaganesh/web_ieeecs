// Sample movie data (you can replace this with actual API calls)
const movies = [
    { id: 1, title: 'Movie 1', image: 'https://via.placeholder.com/150', overview: 'Overview of Movie 1' },
    { id: 2, title: 'Movie 2', image: 'https://via.placeholder.com/150', overview: 'Overview of Movie 2' },
    { id: 3, title: 'Movie 3', image: 'https://via.placeholder.com/150', overview: 'Overview of Movie 3' }
];

// Function to display movies on the page
function displayMovies() {
    const movieGrid = document.getElementById('movieGrid');

    movies.forEach(movie => {
        const card = document.createElement('div');
        card.classList.add('card');
        card.innerHTML = `
            <img src="${movie.image}" alt="${movie.title}">
            <div class="card-content">
                <h3>${movie.title}</h3>
                <p>${movie.overview}</p>
                <button onclick="addToWatchlist(${movie.id})">Add to Watchlist</button>
            </div>
        `;
        movieGrid.appendChild(card);
    });
}

// Function to add a movie to the watchlist
function addToWatchlist(movieId) {
    const movie = movies.find(m => m.id === movieId);
    // Add movie to watchlist (you can use localStorage or a database for this)
}

// Show login modal
document.getElementById('loginBtn').addEventListener('click', () => {
    document.getElementById('loginModal').style.display = 'block';
});

// Show register modal
document.getElementById('registerBtn').addEventListener('click', () => {
    document.getElementById('registerModal').style.display = 'block';
});

// Close modals
document.querySelectorAll('.close').
