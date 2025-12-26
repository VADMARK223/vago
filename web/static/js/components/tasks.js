(() => {
    window.deleteTask = async function(id) {
        if (!confirm('Удалить задачу?')) return;

        const res = await fetch(`/tasks/${id}`, { method: 'DELETE' });
        if (res.ok) {
            document.getElementById(`task-${id}`).remove();
        } else {
            alert('Error deleting task');
        }
    }

    window.toggleCompleted = async function (id, completed) {
        const res = await fetch(`/tasks/${id}`, {
            method: 'PUT',
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ completed })
        });

        if (!res.ok) {
            const data = await res.json().catch(() => ({}));
            alert("Error task updating: " + (data.error || res.statusText));
            return;
        }

        // Меняем стиль строки без обновления страницы
        const li = document.getElementById(`task-${id}`);
        if (completed) {
            li.classList.add("completed");
        } else {
            li.classList.remove("completed");
        }
    }
})();