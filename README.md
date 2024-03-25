Go Note app

I journal daily using md files.

They are named likes this `YYYY-MM-DD.md`

The goal is to add a more context when having multiple files per day.

We leverage OpenAI API for the job.
We provide the name and the content of the file, with a custom prompt. The bot shall answer with a proper name.


Work in Process : Add suggestions, user validations, and renaming of the files.
Add a pattern to target only the new files.
Maybe by doing a git shell command, and looking for the tracked files.


---

Related Documentation:

[Open AI Model Documentation](https://platform.openai.com/docs/deprecations/2023-11-06-chat-model-updates) 


[Example of API Usage](https://pkg.go.dev/github.com/sashabaranov/go-openai#section-readme)

[Api request model](https://platform.openai.com/docs/quickstart/step-3-sending-your-first-api-request)

[Wireframes](https://www.tldraw.com/v/ASeBNdC0ePsvQQcKzK4Sf?v=-105,-168,3140,1852&p=page)