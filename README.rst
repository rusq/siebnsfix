===============
Siebel NS Fixer
===============
Little tool to fix the Oracle Siebel CRM Gateway Naming Server ``siebns.dat``
file after making manual modifications to it.

This program implements instructions written in this guide_.

.. _guide: https://github.com/rusq/siebns_size/tree/master/guide

Download siebns fix
-------------------

You can download a compiled version for your platform from `Releases page`_.

.. _`Releases page`: https://github.com/rusq/siebnsfix/releases


Compiling from source
---------------------

To compile from source, assuming you have cloned this repo to
``$HOME/src/siebnsfix/``::

  $ cd $HOME/src/siebnsfix
  $ go get github.com/rusq/siebns
  $ go build

  $ ./siebnsfix
  Siebnsfix - fix checksum in Siebel Gateway file

  Usage: siebnsfix <siebns.dat>

Usage
-----
Get your ``siebns.dat`` file that needs correction, and run siebnsfix on it::

  $ ./siebnsfix siebns.dat
  Siebnsfix 1.0.0 - fix checksum in Siebel Gateway file
  2018/01/15 20:59:18 File siebns.dat:  Correction needed.
  2018/01/15 20:59:18 File siebns.dat:  OK: Updated 12 bytes.

