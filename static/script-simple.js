// Упрощенный JavaScript - основная логика перенесена в Go
document.addEventListener('DOMContentLoaded', function() {
    const TASKS_URL = "/api/tasks";
    const AUTH_URL  = "/api/auth";
    const TEAM_PROJECTS_URL = "/api/team-projects";

    // AUTH HELPERS 
    function setToken(token) {
        localStorage.setItem("token", token);
    }
    
    function getToken() {
        return localStorage.getItem("token");
    }
    
    function clearToken() {
        localStorage.removeItem("token");
    }
    
    function getHeaders(extra = {}) {
        const currentLanguage = localStorage.getItem('language') || 'ru';
        return { 
            ...extra, 
            'Accept-Language': currentLanguage 
        };
    }

    function authHeaders(extra = {}) {
        const token = getToken();
        const headers = getHeaders(extra);
        return token
            ? { ...headers, Authorization: `Bearer ${token}` }
            : headers;
    }

    // TRANSLATION HELPER 
    function getTranslation(key, fallback) {
        if (typeof t === 'function') {
            const translation = t(key);
            if (translation && translation !== key) {
                return translation;
            }
        }
        return fallback || key;
    }

    // DEADLINE FORMATTING FUNCTIONS
    function formatDeadlineText(deadline) {
        if (!deadline || !deadline.Valid) {
            return getTranslation('noDeadline', 'Без дедлайна');
        }

        const date = new Date(deadline.Time);
        const now = new Date();
        const diffMs = date.getTime() - now.getTime();
        const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

        const dateStr = date.toLocaleDateString('ru-RU', { day: '2-digit', month: '2-digit', year: 'numeric' });
        const timeStr = date.toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit' });

        if (diffDays < 0) {
            return getTranslation('overdue', 'Просрочено') + ': ' + dateStr + ' ' + timeStr;
        } else if (diffDays === 0) {
            return getTranslation('today', 'Сегодня') + ': ' + timeStr;
        } else if (diffDays === 1) {
            return getTranslation('tomorrow', 'Завтра') + ': ' + timeStr;
        } else if (diffDays <= 7) {
            const daysText = getTranslation('days', 'дн.');
            return getTranslation('inDays', 'Через') + ' ' + diffDays + ' ' + daysText + ': ' + dateStr + ' ' + timeStr;
        } else {
            return getTranslation('until', 'До') + ' ' + dateStr + ' ' + timeStr;
        }
    }

    function getDeadlineClass(deadline, isCompleted) {
        if (!deadline || !deadline.Valid) {
            return 'no-deadline';
        }

        const date = new Date(deadline.Time);
        const now = new Date();

        if (isCompleted) {
            if (date >= now) {
                return 'completed';
            }
            return 'overdue';
        } else {
            if (date < now) {
                return 'overdue';
            }
            return 'upcoming';
        }
    }

    function updateDeadlineTexts() {
        // Обновляем тексты дедлайнов для личных задач
        const taskDeadlines = document.querySelectorAll('#task-list .task-deadline-text');
        taskDeadlines.forEach(element => {
            const taskItem = element.closest('.task-item');
            const taskId = taskItem.dataset.taskId;
            // Здесь нужно получить данные задачи, но для простоты обновим только текст
            // В реальном приложении нужно будет перезагрузить задачи
        });

        // Обновляем тексты дедлайнов для командных задач
        const teamTaskDeadlines = document.querySelectorAll('#team-task-list .task-deadline-text');
        teamTaskDeadlines.forEach(element => {
            const taskItem = element.closest('.team-task-item');
            const taskId = taskItem.dataset.taskId;
            // Аналогично для командных задач
        });

        // Обновляем тексты авторов
        const authorTexts = document.querySelectorAll('#team-task-list .task-author');
        authorTexts.forEach(element => {
            const text = element.textContent;
            const authorName = text.split(': ')[1];
            element.textContent = getTranslation('author', 'Автор') + ': ' + authorName;
        });
    }

    // NOTIFICATION SYSTEM
    function showNotification(message, type = 'info', duration = 3000) {
        const container = document.getElementById('notification-container');
        if (!container) return;

        const notification = document.createElement('div');
        notification.className = `notification ${type}`;
        
        const icons = {
            success: '✅',
            error: '❌',
            warning: '⚠️',
            info: 'ℹ️'
        };
        
        notification.innerHTML = `
            <div class="notification-content">
                <span class="notification-icon">${icons[type] || icons.info}</span>
                <span class="notification-message">${message}</span>
                <button class="notification-close" onclick="closeNotification(this)">×</button>
            </div>
            <div class="notification-progress">
                <div class="notification-progress-bar"></div>
            </div>
        `;
        
        container.appendChild(notification);
        
        setTimeout(() => {
            notification.classList.add('show');
        }, 100);
        
        setTimeout(() => {
            closeNotification(notification.querySelector('.notification-close'));
        }, duration);
    }
    
    function closeNotification(closeBtn) {
        const notification = closeBtn.closest('.notification');
        if (!notification) return;
        
        notification.classList.remove('show');
        setTimeout(() => {
            if (notification.parentNode) {
                notification.parentNode.removeChild(notification);
            }
        }, 300);
    }

    // TASKS - загружаем JSON и форматируем на клиенте
    async function loadTasks() {
        const taskList = document.getElementById("task-list");
        if (!taskList) return;
        
        try {
            const res = await fetch(`${TASKS_URL}`, { headers: authHeaders() });
            if (!res.ok) {
                const errorText = await res.text();
                showNotification(errorText, 'error');
                return;
            }
            
            const tasks = await res.json();
            taskList.innerHTML = '';
            
            tasks.forEach(task => {
                const li = document.createElement('li');
                li.className = 'task-item';
                li.setAttribute('data-task-id', task.id);
                
                const deadlineText = formatDeadlineText(task.deadline);
                const deadlineClass = getDeadlineClass(task.deadline, task.done);
                
                li.innerHTML = `
                    <div class="task-content">
                        <span class="task-title${task.done ? ' done' : ''}">${task.title}</span>
                        <div class="task-deadline-text ${deadlineClass}">${deadlineText}</div>
                    </div>
                    <div class="task-buttons">
                        <button onclick="toggleTask(${task.id}, ${task.done})" class="toggle-btn">
                            ${task.done ? getTranslation('markIncomplete', 'Отметить невыполненной') : getTranslation('markCompleted', 'Отметить выполненной')}
                        </button>
                        <button onclick="editTask(${task.id})" class="edit-btn">${getTranslation('edit', 'Редактировать')}</button>
                        <button onclick="deleteTask(${task.id})" class="delete-btn">${getTranslation('deleteTask', 'Удалить')}</button>
                    </div>
                `;
                taskList.appendChild(li);
            });
            
            // Обновляем кнопки после загрузки
            if (typeof window.updateTaskButtons === 'function') {
                window.updateTaskButtons();
            }
        } catch (error) {
            showNotification('Ошибка загрузки задач', 'error');
        }
    }

    // AUTH FUNCTIONS
    async function register() {
        try {
            const res = await fetch(`${AUTH_URL}/register`, {
                method: "POST",
                headers: getHeaders({ "Content-Type": "application/json" }),
                body: JSON.stringify({
                    username: document.getElementById("reg-username").value.trim(),
                    email: document.getElementById("reg-email").value.trim(),
                    password: document.getElementById("reg-password").value
                })
            });
            
            if (res.ok) {
                const data = await res.json();
                showNotification(data.message || getTranslation('registerSuccess', 'Регистрация успешна!'), 'success');
            } else {
                const errorText = await res.text();
                showNotification(errorText, 'error');
            }
        } catch (error) {
            showNotification('Ошибка регистрации', 'error');
        }
    }

    async function login() {
        try {
            const res = await fetch(`${AUTH_URL}/login`, {
                method: "POST",
                headers: getHeaders({ "Content-Type": "application/json" }),
                body: JSON.stringify({
                    username: document.getElementById("login-username").value.trim(),
                    password: document.getElementById("login-password").value
                })
            });

            if (res.ok) {
                const data = await res.json();
                setToken(data.token);
                document.getElementById("auth").style.display = "none";
                document.getElementById("app").style.display = "block";
                loadTasks();
                showNotification(getTranslation('loginSuccess', 'Успешный вход'), 'success');
            } else {
                const msg = await res.text();
                showNotification(msg, 'error');
            }
        } catch (error) {
            showNotification('Ошибка входа', 'error');
        }
    }

    function logout() {
        clearToken();
        document.getElementById("auth").style.display = "block";
        document.getElementById("app").style.display = "none";
        showNotification(getTranslation('logoutSuccess', 'Вы вышли из системы'), 'info');
    }

    // TASK OPERATIONS
    async function toggleTask(id, currentStatus) {
        try {
            const res = await fetch(`${TASKS_URL}/${id}`, {
                method: "PATCH",
                headers: authHeaders({ "Content-Type": "application/json" }),
                body: JSON.stringify({ done: !currentStatus }),
            });
            
            if (res.ok) {
                loadTasks(); // Перезагружаем список
            } else {
                showNotification('Ошибка обновления задачи', 'error');
            }
        } catch (error) {
            showNotification('Ошибка обновления задачи', 'error');
        }
    }

    async function deleteTask(id) {
        try {
            const res = await fetch(`${TASKS_URL}/${id}`, { 
                method: "DELETE", 
                headers: authHeaders() 
            });
            
            if (res.ok) {
                loadTasks(); // Перезагружаем список
            } else {
                showNotification('Ошибка удаления задачи', 'error');
            }
        } catch (error) {
            showNotification('Ошибка удаления задачи', 'error');
        }
    }

    async function editTask(id) {
        // Получаем данные задачи из DOM
        const taskElement = document.querySelector(`[data-task-id="${id}"]`);
        if (!taskElement) return;
        
        const titleElement = taskElement.querySelector('.task-title');
        const deadlineElement = taskElement.querySelector('.task-deadline-text');
        
        if (!titleElement) return;
        
        const currentTitle = titleElement.textContent;
        const currentDeadline = deadlineElement ? deadlineElement.textContent : '';
        
        // Создаем форму редактирования
        const editForm = `
            <div class="edit-form">
                <input type="text" class="edit-title" value="${currentTitle}" placeholder="Название задачи">
                <input type="datetime-local" class="edit-deadline" placeholder="Дедлайн">
                <button onclick="saveTask(${id})" class="save-btn">${getTranslation('save', 'Сохранить')}</button>
                <button onclick="cancelEdit(${id})" class="cancel-btn">${getTranslation('cancel', 'Отмена')}</button>
            </div>
        `;
        
        taskElement.innerHTML = editForm;
    }

    async function saveTask(id) {
        const titleInput = document.querySelector(`[data-task-id="${id}"] .edit-title`);
        const deadlineInput = document.querySelector(`[data-task-id="${id}"] .edit-deadline`);
        
        if (!titleInput || !deadlineInput) return;
        
        const title = titleInput.value.trim();
        if (!title) {
            showNotification('Название задачи не может быть пустым!', 'warning');
            return;
        }
        
        const deadline = deadlineInput.value ? new Date(deadlineInput.value).toISOString() : null;
        
        try {
            const res = await fetch(`${TASKS_URL}/${id}`, {
                method: "PATCH",
                headers: authHeaders({ "Content-Type": "application/json" }),
                body: JSON.stringify({ title, deadline }),
            });
            
            if (res.ok) {
                loadTasks(); // Перезагружаем список
            } else {
                showNotification('Ошибка сохранения задачи', 'error');
            }
        } catch (error) {
            showNotification('Ошибка сохранения задачи', 'error');
        }
    }

    function cancelEdit(id) {
        loadTasks(); // Перезагружаем список
    }

    // ADD TASK
    async function addTask() {
        const title = document.getElementById("task-input").value.trim();
        if (!title) return;
        
        const deadlineInput = document.getElementById('task-deadline');
        const deadline = deadlineInput.value ? new Date(deadlineInput.value).toISOString() : null;
        
        try {
            const res = await fetch(TASKS_URL, {
                method: "POST",
                headers: authHeaders({ "Content-Type": "application/json" }),
                body: JSON.stringify({ title, deadline }),
            });
            
            if (res.ok) {
                document.getElementById("task-input").value = "";
                deadlineInput.value = "";
                loadTasks();
                showNotification(getTranslation('taskAdded', 'Задача добавлена'), 'success');
            } else {
                const errorText = await res.text();
                showNotification(errorText, 'error');
            }
        } catch (error) {
            showNotification('Ошибка добавления задачи', 'error');
        }
    }

    // PASSWORD RESET
    function showForgotPassword() {
        document.getElementById("forgot-password").style.display = "block";
    }

    function hideForgotPassword() {
        document.getElementById("forgot-password").style.display = "none";
    }

    async function requestPasswordReset() {
        const email = document.getElementById("forgot-email").value.trim();
        if (!email) {
            showNotification(getTranslation('emailRequired', 'Требуется email'), 'warning');
            return;
        }

        try {
            const res = await fetch(`${AUTH_URL}/request-password-reset`, {
                method: "POST",
                headers: getHeaders({ "Content-Type": "application/json" }),
                body: JSON.stringify({ email }),
            });

            if (res.ok) {
                const data = await res.json();
                showNotification(data.message || getTranslation('passwordResetSent', 'Ссылка для сброса пароля отправлена'), 'success');
                hideForgotPassword();
            } else {
                const errorText = await res.text();
                showNotification(errorText, 'error');
            }
        } catch (error) {
            showNotification('Ошибка сброса пароля', 'error');
        }
    }

    // TEAM PROJECTS (упрощенная версия)
    let currentProject = null;

    async function showMyTasks() {
        document.getElementById('my-tasks-section').style.display = 'block';
        document.getElementById('team-projects-section').style.display = 'none';
        document.getElementById('create-project-section').style.display = 'none';
        document.getElementById('join-project-section').style.display = 'none';
        document.getElementById('team-project-section').style.display = 'none';
        
        document.getElementById('my-tasks-btn').classList.add('active');
        document.getElementById('team-projects-btn').classList.remove('active');
        
        loadTasks();
    }

    async function showTeamProjects() {
        document.getElementById('my-tasks-section').style.display = 'none';
        document.getElementById('team-projects-section').style.display = 'block';
        document.getElementById('create-project-section').style.display = 'none';
        document.getElementById('join-project-section').style.display = 'none';
        document.getElementById('team-project-section').style.display = 'none';
        
        document.getElementById('my-tasks-btn').classList.remove('active');
        document.getElementById('team-projects-btn').classList.add('active');
        
        loadTeamProjects();
    }

    async function loadTeamProjects() {
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/my`, { headers: authHeaders() });
            if (!res.ok) {
                const errorText = await res.text();
                showNotification(errorText, 'error');
                return;
            }
            
            const projects = await res.json();
            const projectsList = document.getElementById('team-projects-list');
            projectsList.innerHTML = '';
            
            if (projects.length === 0) {
                projectsList.innerHTML = `<p>${getTranslation('noProjects', 'У вас пока нет командных проектов')}</p>`;
                return;
            }
            
            projects.forEach(project => {
                const div = document.createElement('div');
                div.className = 'team-project-item';
                div.innerHTML = `
                    <span>${project.name}</span>
                    <small>${getTranslation('code', 'Код')}: ${project.code}</small>
                    <button onclick="openTeamProject(${project.id}, '${project.name}')">${getTranslation('open', 'Открыть')}</button>
                `;
                projectsList.appendChild(div);
            });
        } catch (error) {
            showNotification('Ошибка загрузки проектов', 'error');
        }
    }

    // Остальные функции командных проектов (упрощенные)
    function showCreateProject() {
        document.getElementById('team-projects-section').style.display = 'none';
        document.getElementById('create-project-section').style.display = 'block';
        document.getElementById('project-name').value = '';
    }

    function showJoinProject() {
        document.getElementById('team-projects-section').style.display = 'none';
        document.getElementById('join-project-section').style.display = 'block';
        document.getElementById('project-code').value = '';
    }

    async function createProject() {
        const name = document.getElementById('project-name').value.trim();
        if (!name) {
            showNotification(getTranslation('fieldRequired', 'Поле обязательно для заполнения'), 'warning');
            return;
        }
        
        try {
            const res = await fetch(TEAM_PROJECTS_URL, {
                method: 'POST',
                headers: authHeaders({ 'Content-Type': 'application/json' }),
                body: JSON.stringify({ name })
            });
            
            if (res.ok) {
                const project = await res.json();
                showNotification(getTranslation('projectCreated', 'Проект создан! Код: {code}').replace('{code}', project.code), 'success');
                showTeamProjects();
            } else {
                const errorText = await res.text();
                showNotification(errorText, 'error');
            }
        } catch (error) {
            showNotification('Ошибка создания проекта', 'error');
        }
    }

    async function joinProject() {
        const code = document.getElementById('project-code').value.trim();
        if (!code) {
            showNotification(getTranslation('fieldRequired', 'Поле обязательно для заполнения'), 'warning');
            return;
        }
        
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/join`, {
                method: 'POST',
                headers: authHeaders({ 'Content-Type': 'application/json' }),
                body: JSON.stringify({ code })
            });
            
            if (res.ok) {
                showNotification(getTranslation('projectJoined', 'Вы присоединились к проекту'), 'success');
                showTeamProjects();
            } else {
                const errorText = await res.text();
                showNotification(errorText, 'error');
            }
        } catch (error) {
            showNotification('Ошибка присоединения к проекту', 'error');
        }
    }

    async function openTeamProject(projectId, projectName) {
        currentProject = { id: projectId, name: projectName };
        
        document.getElementById('team-projects-section').style.display = 'none';
        document.getElementById('team-project-section').style.display = 'block';
        document.getElementById('project-info').textContent = `${projectName} (ID: ${projectId})`;
        
        loadTeamTasks(projectId);
    }

    async function loadTeamTasks(projectId) {
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/${projectId}/tasks`, { headers: authHeaders() });
            if (!res.ok) {
                const errorText = await res.text();
                showNotification(errorText, 'error');
                return;
            }
            
            const tasks = await res.json();
            const taskList = document.getElementById('team-task-list');
            taskList.innerHTML = '';
            
            tasks.forEach(task => {
                const li = document.createElement('li');
                li.className = 'team-task-item';
                li.setAttribute('data-task-id', task.id);
                
                const deadlineText = formatDeadlineText(task.deadline);
                const deadlineClass = getDeadlineClass(task.deadline, task.done);
                const authorText = getTranslation('author', 'Автор');
                
                li.innerHTML = `
                    <div class="task-content">
                        <span class="task-title${task.done ? ' done' : ''}">${task.title}</span>
                        <div class="task-deadline-text ${deadlineClass}">${deadlineText}</div>
                        ${task.created_by_username ? `<div class="task-author">${authorText}: ${task.created_by_username}</div>` : ''}
                    </div>
                    <div class="task-buttons">
                        <button onclick="toggleTeamTask(${projectId}, ${task.id}, ${task.done})" class="toggle-btn">
                            ${task.done ? getTranslation('markIncomplete', 'Отметить невыполненной') : getTranslation('markCompleted', 'Отметить выполненной')}
                        </button>
                        <button onclick="editTeamTask(${projectId}, ${task.id})" class="edit-btn">${getTranslation('edit', 'Редактировать')}</button>
                        <button onclick="deleteTeamTask(${projectId}, ${task.id})" class="delete-btn">${getTranslation('deleteTask', 'Удалить')}</button>
                    </div>
                `;
                taskList.appendChild(li);
            });
            
            // Обновляем кнопки после загрузки
            if (typeof window.updateTaskButtons === 'function') {
                window.updateTaskButtons();
            }
        } catch (error) {
            showNotification('Ошибка загрузки командных задач', 'error');
        }
    }

    // Простые функции для командных задач
    async function toggleTeamTask(projectId, taskId, currentStatus) {
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/${projectId}/tasks/${taskId}`, {
                method: 'PATCH',
                headers: authHeaders({ 'Content-Type': 'application/json' }),
                body: JSON.stringify({ done: !currentStatus }),
            });
            
            if (res.ok) {
                loadTeamTasks(projectId);
            } else {
                showNotification('Ошибка обновления задачи', 'error');
            }
        } catch (error) {
            showNotification('Ошибка обновления задачи', 'error');
        }
    }

    async function editTeamTask(projectId, taskId) {
        // Получаем данные задачи из DOM
        const taskElement = document.querySelector(`[data-task-id="${taskId}"]`);
        if (!taskElement) return;
        
        const titleElement = taskElement.querySelector('.task-title');
        const deadlineElement = taskElement.querySelector('.task-deadline-text');
        
        if (!titleElement) return;
        
        const currentTitle = titleElement.textContent;
        const currentDeadline = deadlineElement ? deadlineElement.textContent : '';
        
        // Создаем форму редактирования
        const editForm = `
            <div class="edit-form">
                <input type="text" class="edit-title" value="${currentTitle}" placeholder="Название задачи">
                <input type="datetime-local" class="edit-deadline" placeholder="Дедлайн">
                <button onclick="saveTeamTask(${projectId}, ${taskId})" class="save-btn">${getTranslation('save', 'Сохранить')}</button>
                <button onclick="cancelTeamEdit(${projectId}, ${taskId})" class="cancel-btn">${getTranslation('cancel', 'Отмена')}</button>
            </div>
        `;
        
        taskElement.innerHTML = editForm;
    }

    async function saveTeamTask(projectId, taskId) {
        const titleInput = document.querySelector(`[data-task-id="${taskId}"] .edit-title`);
        const deadlineInput = document.querySelector(`[data-task-id="${taskId}"] .edit-deadline`);
        
        if (!titleInput || !deadlineInput) return;
        
        const title = titleInput.value.trim();
        if (!title) {
            showNotification('Название задачи не может быть пустым!', 'warning');
            return;
        }
        
        const deadline = deadlineInput.value ? new Date(deadlineInput.value).toISOString() : null;
        
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/${projectId}/tasks/${taskId}`, {
                method: 'PATCH',
                headers: authHeaders({ 'Content-Type': 'application/json' }),
                body: JSON.stringify({ title, deadline }),
            });
            
            if (res.ok) {
                loadTeamTasks(projectId); // Перезагружаем список
            } else {
                showNotification('Ошибка сохранения задачи', 'error');
            }
        } catch (error) {
            showNotification('Ошибка сохранения задачи', 'error');
        }
    }

    function cancelTeamEdit(projectId, taskId) {
        loadTeamTasks(projectId); // Перезагружаем список
    }

    async function deleteTeamTask(projectId, taskId) {
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/${projectId}/tasks/${taskId}`, { 
                method: 'DELETE', 
                headers: authHeaders() 
            });
            
            if (res.ok) {
                loadTeamTasks(projectId);
            } else {
                showNotification('Ошибка удаления задачи', 'error');
            }
        } catch (error) {
            showNotification('Ошибка удаления задачи', 'error');
        }
    }
    
    // ADD TEAM TASK
    async function addTeamTask() {
        const title = document.getElementById('team-task-input').value.trim();
        if (!title || !currentProject) return;
        
        const deadlineInput = document.getElementById('team-task-deadline');
        const deadline = deadlineInput.value ? new Date(deadlineInput.value).toISOString() : null;
        
        try {
            const res = await fetch(`${TEAM_PROJECTS_URL}/${currentProject.id}/tasks`, {
                method: 'POST',
                headers: authHeaders({ 'Content-Type': 'application/json' }),
                body: JSON.stringify({ title, deadline }),
            });
            
            if (res.ok) {
                document.getElementById('team-task-input').value = '';
                deadlineInput.value = '';
                loadTeamTasks(currentProject.id);
            } else {
                const errorText = await res.text();
                showNotification(errorText, 'error');
            }
        } catch (error) {
            showNotification('Ошибка добавления задачи', 'error');
        }
    }

    // EVENT LISTENERS
    document.getElementById("add-btn").onclick = addTask;
    document.getElementById('add-team-task-btn').onclick = addTeamTask;

    // Глобальные функции для HTML
    window.register = register;
    window.login = login;
    window.logout = logout;
    window.showForgotPassword = showForgotPassword;
    window.hideForgotPassword = hideForgotPassword;
    window.requestPasswordReset = requestPasswordReset;
    window.showMyTasks = showMyTasks;
    window.showTeamProjects = showTeamProjects;
    window.showCreateProject = showCreateProject;
    window.showJoinProject = showJoinProject;
    window.createProject = createProject;
    window.joinProject = joinProject;
    window.openTeamProject = openTeamProject;
    window.showNotification = showNotification;
    window.closeNotification = closeNotification;
    window.toggleTask = toggleTask;
    window.deleteTask = deleteTask;
    window.editTask = editTask;
    window.saveTask = saveTask;
    window.cancelEdit = cancelEdit;
    window.toggleTeamTask = toggleTeamTask;
    window.editTeamTask = editTeamTask;
    window.saveTeamTask = saveTeamTask;
    window.cancelTeamEdit = cancelTeamEdit;
    window.deleteTeamTask = deleteTeamTask;
    window.loadTasks = loadTasks;
    
    // Проверяем, есть ли сохраненный токен при загрузке страницы
    const token = getToken();
    if (token) {
        document.getElementById("auth").style.display = "none";
        document.getElementById("app").style.display = "block";
        loadTasks();
    }
});
