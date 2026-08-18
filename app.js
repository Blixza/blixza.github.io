document.addEventListener("DOMContentLoaded", () => {
    fetchLibraryData();
});

async function fetchLibraryData() {
    const statusElement = document.getElementById("status-message");
    const tbody = document.getElementById("album-body");

    try {
        const response = await fetch("library.json");
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const records = await response.json();

        if (!records || records.length === 0) {
            statusElement.textContent = "No album records found in library.json.";
            return;
        }

        statusElement.style.display = "none";

        records.forEach(record => {
            const row = createTableRow(record);
            tbody.appendChild(row);
        });

    } catch (error) {
        console.error("Error loading library.json:", error);
        statusElement.textContent = "Failed to load library data. Ensure library.json contains valid JSON.";
    }
}

function createTableRow(record) {
    const tr = document.createElement("tr");

    const artistsStr = Array.isArray(record.artist) ? record.artist.join(", ") : record.artist || "—";

    const formatSongList = (songs) => {
        if (!songs || songs.length === 0) return "—";
        return `<ul class="song-list">${songs.map(song => `<li>• ${escapeHtml(song)}</li>`).join("")}</ul>`;
    };

    const isSaved = String(record.saved).toLowerCase() === "true" || String(record.saved).toLowerCase() === "yes";
    const savedBadge = isSaved
        ? `<span class="saved-badge">Yes</span>`
        : `<span class="saved-badge no">No</span>`;

    tr.innerHTML = `
    <td>${escapeHtml(record.date || "—")}</td>
    <td><strong>${escapeHtml(artistsStr)}</strong></td>
    <td>${escapeHtml(record.album || "—")}</td>
    <td class="rating">${escapeHtml(record.rating || "—")}</td>
    <td>${savedBadge}</td>
    <td>${formatSongList(record.best_songs)}</td>
    <td>${formatSongList(record.worst_songs)}</td>
    <td>${escapeHtml(record.thoughts || "—")}</td>
  `;

    return tr;
}

function escapeHtml(str) {
    if (typeof str !== "string") return str;
    return str
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#039;");
}