% Installing and configuring the Calamares installer

(some rambling)

Four years ago, in 2021 and then in 2023, I posted two very low quality videos
on YouTube on installing and configuring Calamares, which gained a lot of
traction. I was of course quite surprised that people really couldn't find
anything better than such low quality videos in 2021. However, considering the
lack of any proper guide to use Calamares, it made a bit more sense, and I am
extremely bored, so here I am writing about how you can setup Calamares, and
might make another video about it, with better quality, scripting, and editing.

(the actual content)

One of the misunderstandings people have about Calamares is that they think
that it has to somehow become part of their distribution's installation ISO. On
the contrary, it is simply a program that many distributions (Manjaro mainly
comes to mind) runs at the startup on their installation environment.
Throughout this post I will be considering Arch based distributions as Woof OS
is based on Arch and my experience comes solely from its development.

I think that this explanation fixes the confusion a lot of people face during
their first attempt at making their own personal respin of Arch. Since it is
just another program, you configure it just like any other programs. The only
difference is that you run it on startup in your installation ISO and not on
the installed environment itself (this can, for example, be achieved by adding
Calamares to startup in the home directory of your installation user in
archlive and _not_ including it in /etc/skel).

The AUR now includes a regularly updated Calamres package, so to include
Calamares in your build of archiso, you would add that AUR package to your own
repository (See [Packages in the Woof OS Pacman
Repository](https://os.lewoof.xyz/packages)).
