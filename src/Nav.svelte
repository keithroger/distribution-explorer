
<nav class="navbar-container">
    <div class="navbar-header">
        <a href="/">DISTRIBUTION EXPLORER</a>
    </div>
    <ul class=navbar-menu-items>
        <li>
            <a href="/normal">Normal</a> 
        </li>
        <li>
            <a href="/uniform">Uniform</a>
        </li>
        <li>
            <a href="/binomial">Binomial</a>
        </li>
        <li>
            <a href="/poisson">Poisson</a>
        </li>
        <li>
            <a href="/gamma">Gamma</a>
        </li>
        <li>
            <a href="/beta">Beta</a>
        </li>
    </ul>
</nav>

<style>
    .navbar-container {
        display: flex;
        flex-direction: column;
        margin: 0.5em;
    }

    .navbar-header {
        text-align: center;
        margin: 0.25em;
        font-size: 2em;
    }

    ul {
        list-style-type: none;
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        margin: 0;
        padding: 0;
    }
    li {
        max-width: 130px;
        min-width: 60px;
    }

    li a {
        display: block;
        padding: 8px 16px;
        text-decoration: none;
    }

</style>