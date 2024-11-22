<img src="https://github.com/woofslides/.github/blob/main/profile/logo.png?raw=true" width="100" height="100">

# Presentations with HTML, and more.

It starts with Linux, and me having to make presentations. I don't use office suites, just like a stereotypical Arch user.
And when I had to write documents and make presentations, well, I was lost.

For documents, I've been using LaTeX; it works pretty well, and I like it.

For presentations, I've tried Google Slides and Beamer, neither of which I liked using.
I did realise however, that I liked using Beamer more, not only because I could write it in plain text and I can use whatever text editor I wanted to, but mainly because it helped me focus more on what the content of the presentation was before I considered how the presentation would look.
Now this is a double-edged sword, while focusing exclusively on the content of the rather than its styling is advantagous for documents (which made me enjoy using LaTeX), I couldn't say the same about presentations.
While I do think that Beamer was better than the most used presentation applications like Microsoft Powerpoint and Google Slides because of its focus on the content, presentations are meant to be presented, and the presentation of a presentation is as important as its content - it's called a _presentation_.

We now have three requirements for how the presentation is made:

- The main focus should be on the content.
- The presentation of the presentation should have as much importance during the writing of the presentation as much as the content itself, but it should not obstruct the user's focus on the content.
- The source files should be written in plain text files that can be edited using any text editor.

Markup languages aim to meet all of these requirements, and can generally be compiled into one another.
Now there are apps to make presentations using Markdown, but there still remains another issue.

During a presentation, you tend to connect your device to a projector and then use the stage as your space to explain.
I generally tend to not stay near the device when presenting, so having to walk back to change what slide is visible is kinda irritating.
I figured it would be much better if I could use my phone or any handheld device to control the presentation.

We then boil it down to two requirements:

- The source files should be written in a markup language.
- During the presentation, any approved device on the local network should be able to control the presentation.

I haven't done much searching because I'm lazy, but from what I've seen so far, no app does this the way I want it to do, so I'm building my own presentation system made of a viewer, a server, and the controller. I might write more as I work on it.

---

Check out [WoofSlides](https://github.com/woofslides) on GitHub, and my [Notion site](https://acutewoof.notion.site/Woof-slides-11e57c8022a380679f1ce26ee0a480e5) for update notes that aren't made out to be blog posts yet.
