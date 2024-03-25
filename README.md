Go Note app

I journal daily using md files.

They are named likes this `YYYY-MM-DD.md`

The goal is to add a more context when having multiple files per day.

We leverage OpenAI API for the job.
We provide the name and the content of the file, with a custom prompt. The bot shall answer with a proper name.


Work in Process : Add suggestions, user validations, and renaming of the files.
Add a pattern to target only the new files.
Maybe by doing a git shell command, and looking for the tracked files.
